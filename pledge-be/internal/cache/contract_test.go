package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-dev-frame/sponge/pkg/gotest"
	"github.com/go-dev-frame/sponge/pkg/utils"

	"pledge-be/internal/database"
	"pledge-be/internal/model"
)

func newContractCache() *gotest.Cache {
	record1 := &model.Contract{}
	record1.ID = 1
	record2 := &model.Contract{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewContractCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_contractCache_Set(t *testing.T) {
	c := newContractCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Contract)
	err := c.ICache.(ContractCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(ContractCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_contractCache_Get(t *testing.T) {
	c := newContractCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Contract)
	err := c.ICache.(ContractCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(ContractCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(ContractCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_contractCache_MultiGet(t *testing.T) {
	c := newContractCache()
	defer c.Close()

	var testData []*model.Contract
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Contract))
	}

	err := c.ICache.(ContractCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(ContractCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Contract))
	}
}

func Test_contractCache_MultiSet(t *testing.T) {
	c := newContractCache()
	defer c.Close()

	var testData []*model.Contract
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Contract))
	}

	err := c.ICache.(ContractCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_contractCache_Del(t *testing.T) {
	c := newContractCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Contract)
	err := c.ICache.(ContractCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_contractCache_SetCacheWithNotFound(t *testing.T) {
	c := newContractCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Contract)
	err := c.ICache.(ContractCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(ContractCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewContractCache(t *testing.T) {
	c := NewContractCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewContractCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewContractCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
