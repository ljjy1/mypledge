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

func newTokenInfoCache() *gotest.Cache {
	record1 := &model.TokenInfo{}
	record1.ID = 1
	record2 := &model.TokenInfo{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewTokenInfoCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_tokenInfoCache_Set(t *testing.T) {
	c := newTokenInfoCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.TokenInfo)
	err := c.ICache.(TokenInfoCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(TokenInfoCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_tokenInfoCache_Get(t *testing.T) {
	c := newTokenInfoCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.TokenInfo)
	err := c.ICache.(TokenInfoCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(TokenInfoCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(TokenInfoCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_tokenInfoCache_MultiGet(t *testing.T) {
	c := newTokenInfoCache()
	defer c.Close()

	var testData []*model.TokenInfo
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.TokenInfo))
	}

	err := c.ICache.(TokenInfoCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(TokenInfoCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.TokenInfo))
	}
}

func Test_tokenInfoCache_MultiSet(t *testing.T) {
	c := newTokenInfoCache()
	defer c.Close()

	var testData []*model.TokenInfo
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.TokenInfo))
	}

	err := c.ICache.(TokenInfoCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_tokenInfoCache_Del(t *testing.T) {
	c := newTokenInfoCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.TokenInfo)
	err := c.ICache.(TokenInfoCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_tokenInfoCache_SetCacheWithNotFound(t *testing.T) {
	c := newTokenInfoCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.TokenInfo)
	err := c.ICache.(TokenInfoCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(TokenInfoCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewTokenInfoCache(t *testing.T) {
	c := NewTokenInfoCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewTokenInfoCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewTokenInfoCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
