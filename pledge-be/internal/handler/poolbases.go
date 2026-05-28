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

var _ PoolbasesHandler = (*poolbasesHandler)(nil)

// PoolbasesHandler 借贷池基础信息处理器接口
type PoolbasesHandler interface {
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
	UpdateByID(c *gin.Context)
	GetByID(c *gin.Context)
	List(c *gin.Context)
}

// poolbasesHandler 借贷池基础信息数据处理实现
type poolbasesHandler struct {
	iDao dao.PoolbasesDao
}

// NewPoolbasesHandler 创建借贷池基础信息处理器实例
func NewPoolbasesHandler() PoolbasesHandler {
	return &poolbasesHandler{
		iDao: dao.NewPoolbasesDao(
			database.GetDB(), // db driver is mysql
			cache.NewPoolbasesCache(database.GetCacheType()),
		),
	}
}

// Create 创建新的借贷池基础信息记录
// @Summary Create a new poolbases
// @Description Creates a new poolbases entity using the provided data in the request body.
// @Tags poolbases
// @Accept json
// @Produce json
// @Param data body types.CreatePoolbasesRequest true "poolbases information"
// @Success 200 {object} types.CreatePoolbasesReply{}
// @Router /api/v1/poolbases [post]
// @Security BearerAuth
func (h *poolbasesHandler) Create(c *gin.Context) {
	form := &types.CreatePoolbasesRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	poolbases := &model.Poolbases{}
	err = copier.Copy(poolbases, form)
	if err != nil {
		response.Error(c, ecode.ErrCreatePoolbases)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, poolbases)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": poolbases.ID})
}

// DeleteByID 根据 ID 删除借贷池基础信息记录
// @Summary Delete a poolbases by id
// @Description Deletes a existing poolbases identified by the given id in the path.
// @Tags poolbases
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeletePoolbasesByIDReply{}
// @Router /api/v1/poolbases/{id} [delete]
// @Security BearerAuth
func (h *poolbasesHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getPoolbasesIDFromPath(c)
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

// UpdateByID 根据 ID 更新借贷池基础信息记录
// @Summary Update a poolbases by id
// @Description Updates the specified poolbases by given id in the path, support partial update.
// @Tags poolbases
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdatePoolbasesByIDRequest true "poolbases information"
// @Success 200 {object} types.UpdatePoolbasesByIDReply{}
// @Router /api/v1/poolbases/{id} [put]
// @Security BearerAuth
func (h *poolbasesHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getPoolbasesIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdatePoolbasesByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	poolbases := &model.Poolbases{}
	err = copier.Copy(poolbases, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDPoolbases)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, poolbases)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID 根据 ID 查询借贷池基础信息记录
// @Summary Get a poolbases by id
// @Description Gets detailed information of a poolbases specified by the given id in the path.
// @Tags poolbases
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetPoolbasesByIDReply{}
// @Router /api/v1/poolbases/{id} [get]
// @Security BearerAuth
func (h *poolbasesHandler) GetByID(c *gin.Context) {
	_, id, isAbort := getPoolbasesIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	poolbases, err := h.iDao.GetByID(ctx, id)
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

	data := &types.PoolbasesObjDetail{}
	err = copier.Copy(data, poolbases)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDPoolbases)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	response.Success(c, gin.H{"poolbases": data})
}

// List 分页查询借贷池基础信息列表，支持自定义筛选条件
// @Summary Get a paginated list of poolbasess by custom conditions
// @Description Returns a paginated list of poolbases based on query filters, including page number and size.
// @Tags poolbases
// @Accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListPoolbasessReply{}
// @Router /api/v1/poolbases/list [post]
// @Security BearerAuth
func (h *poolbasesHandler) List(c *gin.Context) {
	form := &types.ListPoolbasessRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	poolbasess, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertPoolbasess(poolbasess)
	if err != nil {
		response.Error(c, ecode.ErrListPoolbases)
		return
	}

	response.Success(c, gin.H{
		"poolbasess": data,
		"total":      total,
	})
}

// getPoolbasesIDFromPath 从请求路径中解析借贷池基础信息 ID，返回 (id字符串, id数值, 是否应中止请求)
func getPoolbasesIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

// convertPoolbases 将模型层的 Poolbases 结构体转换为响应层结构体
func convertPoolbases(poolbases *model.Poolbases) (*types.PoolbasesObjDetail, error) {
	data := &types.PoolbasesObjDetail{}
	err := copier.Copy(data, poolbases)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	return data, nil
}

// convertPoolbasess 将模型层的 Poolbases 列表批量转换为响应层结构体
func convertPoolbasess(fromValues []*model.Poolbases) ([]*types.PoolbasesObjDetail, error) {
	toValues := []*types.PoolbasesObjDetail{}
	for _, v := range fromValues {
		data, err := convertPoolbases(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}
