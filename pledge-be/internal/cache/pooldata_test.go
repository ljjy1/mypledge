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

func newPooldataCache() *gotest.Cache {
	record1 := &model.Pooldata{}
	record1.ID = 1
	record2 := &model.Pooldata{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewPooldataCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_pooldataCache_Set(t *testing.T) {
	c := newPooldataCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Pooldata)
	err := c.ICache.(PooldataCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(PooldataCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_pooldataCache_Get(t *testing.T) {
	c := newPooldataCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Pooldata)
	err := c.ICache.(PooldataCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PooldataCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(PooldataCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_pooldataCache_MultiGet(t *testing.T) {
	c := newPooldataCache()
	defer c.Close()

	var testData []*model.Pooldata
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Pooldata))
	}

	err := c.ICache.(PooldataCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PooldataCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Pooldata))
	}
}

func Test_pooldataCache_MultiSet(t *testing.T) {
	c := newPooldataCache()
	defer c.Close()

	var testData []*model.Pooldata
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Pooldata))
	}

	err := c.ICache.(PooldataCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_pooldataCache_Del(t *testing.T) {
	c := newPooldataCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Pooldata)
	err := c.ICache.(PooldataCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_pooldataCache_SetCacheWithNotFound(t *testing.T) {
	c := newPooldataCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Pooldata)
	err := c.ICache.(PooldataCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(PooldataCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewPooldataCache(t *testing.T) {
	c := NewPooldataCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewPooldataCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewPooldataCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
