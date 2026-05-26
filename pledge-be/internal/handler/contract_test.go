package handler

import (
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/go-dev-frame/sponge/pkg/copier"
	"github.com/go-dev-frame/sponge/pkg/gotest"
	"github.com/go-dev-frame/sponge/pkg/httpcli"
	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
	"github.com/go-dev-frame/sponge/pkg/utils"

	"pledge-be/internal/cache"
	"pledge-be/internal/dao"
	"pledge-be/internal/database"
	"pledge-be/internal/model"
	"pledge-be/internal/types"
)

func newContractHandler() *gotest.Handler {
	testData := &model.Contract{}
	testData.ID = 1
	// you can set the other fields of testData here, such as:
	//testData.CreatedAt = time.Now()
	//testData.UpdatedAt = testData.CreatedAt

	// init mock cache
	c := gotest.NewCache(map[string]interface{}{utils.Uint64ToStr(testData.ID): testData})
	c.ICache = cache.NewContractCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})

	// init mock dao
	d := gotest.NewDao(c, testData)
	d.IDao = dao.NewContractDao(d.DB, c.ICache.(cache.ContractCache))

	// init mock handler
	h := gotest.NewHandler(d, testData)
	h.IHandler = &contractHandler{iDao: d.IDao.(dao.ContractDao)}
	iHandler := h.IHandler.(ContractHandler)

	testFns := []gotest.RouterInfo{
		{
			FuncName:    "Create",
			Method:      http.MethodPost,
			Path:        "/contract",
			HandlerFunc: iHandler.Create,
		},
		{
			FuncName:    "DeleteByID",
			Method:      http.MethodDelete,
			Path:        "/contract/:id",
			HandlerFunc: iHandler.DeleteByID,
		},
		{
			FuncName:    "UpdateByID",
			Method:      http.MethodPut,
			Path:        "/contract/:id",
			HandlerFunc: iHandler.UpdateByID,
		},
		{
			FuncName:    "GetByID",
			Method:      http.MethodGet,
			Path:        "/contract/:id",
			HandlerFunc: iHandler.GetByID,
		},
		{
			FuncName:    "List",
			Method:      http.MethodPost,
			Path:        "/contract/list",
			HandlerFunc: iHandler.List,
		},
	}

	h.GoRunHTTPServer(testFns)

	time.Sleep(time.Millisecond * 200)
	return h
}

func Test_contractHandler_Create(t *testing.T) {
	h := newContractHandler()
	defer h.Close()
	testData := &types.CreateContractRequest{}
	_ = copier.Copy(testData, h.TestData.(*model.Contract))

	h.MockDao.SQLMock.ExpectBegin()
	args := h.MockDao.GetAnyArgs(h.TestData)
	h.MockDao.SQLMock.ExpectExec("INSERT INTO .*").
		WithArgs(args[:len(args)-1]...). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(1, 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &httpcli.StdResult{}
	err := httpcli.Post(result, h.GetRequestURL("Create"), testData)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)

}

func Test_contractHandler_DeleteByID(t *testing.T) {
	h := newContractHandler()
	defer h.Close()
	testData := h.TestData.(*model.Contract)
	expectedSQLForDeletion := "DELETE .*"

	h.MockDao.SQLMock.ExpectBegin()
	h.MockDao.SQLMock.ExpectExec(expectedSQLForDeletion).
		WithArgs(testData.ID). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(int64(testData.ID), 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &httpcli.StdResult{}
	err := httpcli.Delete(result, h.GetRequestURL("DeleteByID", testData.ID))
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = httpcli.Delete(result, h.GetRequestURL("DeleteByID", 0))
	assert.NoError(t, err)

	// delete error test
	err = httpcli.Delete(result, h.GetRequestURL("DeleteByID", 111))
	assert.Error(t, err)
}

func Test_contractHandler_UpdateByID(t *testing.T) {
	h := newContractHandler()
	defer h.Close()
	testData := &types.UpdateContractByIDRequest{}
	_ = copier.Copy(testData, h.TestData.(*model.Contract))

	h.MockDao.SQLMock.ExpectBegin()
	h.MockDao.SQLMock.ExpectExec("UPDATE .*").
		WithArgs(h.MockDao.AnyTime, testData.ID). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(int64(testData.ID), 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &httpcli.StdResult{}
	err := httpcli.Put(result, h.GetRequestURL("UpdateByID", testData.ID), testData)
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = httpcli.Put(result, h.GetRequestURL("UpdateByID", 0), testData)
	assert.NoError(t, err)

	// update error test
	err = httpcli.Put(result, h.GetRequestURL("UpdateByID", 111), testData)
	assert.Error(t, err)
}

func Test_contractHandler_GetByID(t *testing.T) {
	h := newContractHandler()
	defer h.Close()
	testData := h.TestData.(*model.Contract)

	// column names and corresponding data
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(testData.ID)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").
		WithArgs(testData.ID, 1).
		WillReturnRows(rows)

	result := &httpcli.StdResult{}
	err := httpcli.Get(result, h.GetRequestURL("GetByID", testData.ID))
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = httpcli.Get(result, h.GetRequestURL("GetByID", 0))
	assert.NoError(t, err)

	// get error test
	err = httpcli.Get(result, h.GetRequestURL("GetByID", 111))
	assert.Error(t, err)
}

func Test_contractHandler_List(t *testing.T) {
	h := newContractHandler()
	defer h.Close()
	testData := h.TestData.(*model.Contract)

	// column names and corresponding data
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(testData.ID)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	result := &httpcli.StdResult{}
	err := httpcli.Post(result, h.GetRequestURL("List"), &types.ListContractsRequest{query.Params{
		Page:  0,
		Limit: 10,
		Sort:  "ignore count", // ignore test count
	}})
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// nil params error test
	err = httpcli.Post(result, h.GetRequestURL("List"), nil)
	assert.NoError(t, err)

	// get error test
	err = httpcli.Post(result, h.GetRequestURL("List"), &types.ListContractsRequest{query.Params{
		Page:  0,
		Limit: 10,
		Sort:  "unknown-column",
	}})
	assert.Error(t, err)
}

func TestNewContractHandler(t *testing.T) {
	defer func() {
		recover()
	}()
	_ = NewContractHandler()
}
