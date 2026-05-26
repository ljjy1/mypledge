package handler

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/go-dev-frame/sponge/pkg/copier"
	"github.com/go-dev-frame/sponge/pkg/gin/middleware"
	"github.com/go-dev-frame/sponge/pkg/gin/response"
	"github.com/go-dev-frame/sponge/pkg/logger"
	"github.com/go-dev-frame/sponge/pkg/utils"

	"pledge-be/internal/cache"
	"pledge-be/internal/dao"
	"pledge-be/internal/database"
	"pledge-be/internal/ecode"
	"pledge-be/internal/model"
	"pledge-be/internal/types"
)

var _ TokenInfoHandler = (*tokenInfoHandler)(nil)

// TokenInfoHandler defining the handler interface
type TokenInfoHandler interface {
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
	UpdateByID(c *gin.Context)
	GetByID(c *gin.Context)
	List(c *gin.Context)
}

type tokenInfoHandler struct {
	iDao dao.TokenInfoDao
}

// NewTokenInfoHandler creating the handler interface
func NewTokenInfoHandler() TokenInfoHandler {
	return &tokenInfoHandler{
		iDao: dao.NewTokenInfoDao(
			database.GetDB(), // db driver is mysql
			cache.NewTokenInfoCache(database.GetCacheType()),
		),
	}
}

// Create a new tokenInfo
// @Summary Create a new tokenInfo
// @Description Creates a new tokenInfo entity using the provided data in the request body.
// @Tags tokenInfo
// @Accept json
// @Produce json
// @Param data body types.CreateTokenInfoRequest true "tokenInfo information"
// @Success 200 {object} types.CreateTokenInfoReply{}
// @Router /api/v1/tokenInfo [post]
// @Security BearerAuth
func (h *tokenInfoHandler) Create(c *gin.Context) {
	form := &types.CreateTokenInfoRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	tokenInfo := &model.TokenInfo{}
	err = copier.Copy(tokenInfo, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateTokenInfo)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, tokenInfo)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": tokenInfo.ID})
}

// DeleteByID delete a tokenInfo by id
// @Summary Delete a tokenInfo by id
// @Description Deletes a existing tokenInfo identified by the given id in the path.
// @Tags tokenInfo
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeleteTokenInfoByIDReply{}
// @Router /api/v1/tokenInfo/{id} [delete]
// @Security BearerAuth
func (h *tokenInfoHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getTokenInfoIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	err := h.iDao.DeleteByID(ctx, id)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// UpdateByID update a tokenInfo by id
// @Summary Update a tokenInfo by id
// @Description Updates the specified tokenInfo by given id in the path, support partial update.
// @Tags tokenInfo
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateTokenInfoByIDRequest true "tokenInfo information"
// @Success 200 {object} types.UpdateTokenInfoByIDReply{}
// @Router /api/v1/tokenInfo/{id} [put]
// @Security BearerAuth
func (h *tokenInfoHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getTokenInfoIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateTokenInfoByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	tokenInfo := &model.TokenInfo{}
	err = copier.Copy(tokenInfo, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDTokenInfo)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, tokenInfo)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID get a tokenInfo by id
// @Summary Get a tokenInfo by id
// @Description Gets detailed information of a tokenInfo specified by the given id in the path.
// @Tags tokenInfo
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetTokenInfoByIDReply{}
// @Router /api/v1/tokenInfo/{id} [get]
// @Security BearerAuth
func (h *tokenInfoHandler) GetByID(c *gin.Context) {
	_, id, isAbort := getTokenInfoIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	tokenInfo, err := h.iDao.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrRecordNotFound) {
			logger.Warn("GetByID not found", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
			response.Error(c, ecode.NotFound)
		} else {
			logger.Error("GetByID error", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
			response.Output(c, ecode.InternalServerError.ToHTTPCode())
		}
		return
	}

	data := &types.TokenInfoObjDetail{}
	err = copier.Copy(data, tokenInfo)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDTokenInfo)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	response.Success(c, gin.H{"tokenInfo": data})
}

// List get a paginated list of tokenInfos by custom conditions
// @Summary Get a paginated list of tokenInfos by custom conditions
// @Description Returns a paginated list of tokenInfo based on query filters, including page number and size.
// @Tags tokenInfo
// @Accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListTokenInfosReply{}
// @Router /api/v1/tokenInfo/list [post]
// @Security BearerAuth
func (h *tokenInfoHandler) List(c *gin.Context) {
	form := &types.ListTokenInfosRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	tokenInfos, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertTokenInfos(tokenInfos)
	if err != nil {
		response.Error(c, ecode.ErrListTokenInfo)
		return
	}

	response.Success(c, gin.H{
		"tokenInfos": data,
		"total":      total,
	})
}

func getTokenInfoIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

func convertTokenInfo(tokenInfo *model.TokenInfo) (*types.TokenInfoObjDetail, error) {
	data := &types.TokenInfoObjDetail{}
	err := copier.Copy(data, tokenInfo)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	return data, nil
}

func convertTokenInfos(fromValues []*model.TokenInfo) ([]*types.TokenInfoObjDetail, error) {
	toValues := []*types.TokenInfoObjDetail{}
	for _, v := range fromValues {
		data, err := convertTokenInfo(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}
