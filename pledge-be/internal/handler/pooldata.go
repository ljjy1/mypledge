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

var _ PooldataHandler = (*pooldataHandler)(nil)

// PooldataHandler 借贷池运行数据处理器接口
type PooldataHandler interface {
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
	UpdateByID(c *gin.Context)
	GetByID(c *gin.Context)
	List(c *gin.Context)
}

// pooldataHandler 借贷池运行数据处理实现
type pooldataHandler struct {
	iDao dao.PooldataDao
}

// NewPooldataHandler 创建借贷池运行数据处理器实例
func NewPooldataHandler() PooldataHandler {
	return &pooldataHandler{
		iDao: dao.NewPooldataDao(
			database.GetDB(), // db driver is mysql
			cache.NewPooldataCache(database.GetCacheType()),
		),
	}
}

// Create 创建新的借贷池运行数据记录
// @Summary Create a new pooldata
// @Description Creates a new pooldata entity using the provided data in the request body.
// @Tags pooldata
// @Accept json
// @Produce json
// @Param data body types.CreatePooldataRequest true "pooldata information"
// @Success 200 {object} types.CreatePooldataReply{}
// @Router /api/v1/pooldata [post]
// @Security BearerAuth
func (h *pooldataHandler) Create(c *gin.Context) {
	form := &types.CreatePooldataRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	pooldata := &model.Pooldata{}
	err = copier.Copy(pooldata, form)
	if err != nil {
		response.Error(c, ecode.ErrCreatePooldata)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, pooldata)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": pooldata.ID})
}

// DeleteByID 根据 ID 删除借贷池运行数据记录
// @Summary Delete a pooldata by id
// @Description Deletes a existing pooldata identified by the given id in the path.
// @Tags pooldata
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeletePooldataByIDReply{}
// @Router /api/v1/pooldata/{id} [delete]
// @Security BearerAuth
func (h *pooldataHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getPooldataIDFromPath(c)
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

// UpdateByID 根据 ID 更新借贷池运行数据记录
// @Summary Update a pooldata by id
// @Description Updates the specified pooldata by given id in the path, support partial update.
// @Tags pooldata
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdatePooldataByIDRequest true "pooldata information"
// @Success 200 {object} types.UpdatePooldataByIDReply{}
// @Router /api/v1/pooldata/{id} [put]
// @Security BearerAuth
func (h *pooldataHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getPooldataIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdatePooldataByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	pooldata := &model.Pooldata{}
	err = copier.Copy(pooldata, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDPooldata)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, pooldata)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID 根据 ID 查询借贷池运行数据记录
// @Summary Get a pooldata by id
// @Description Gets detailed information of a pooldata specified by the given id in the path.
// @Tags pooldata
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetPooldataByIDReply{}
// @Router /api/v1/pooldata/{id} [get]
// @Security BearerAuth
func (h *pooldataHandler) GetByID(c *gin.Context) {
	_, id, isAbort := getPooldataIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	pooldata, err := h.iDao.GetByID(ctx, id)
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

	data := &types.PooldataObjDetail{}
	err = copier.Copy(data, pooldata)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDPooldata)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	response.Success(c, gin.H{"pooldata": data})
}

// List 分页查询借贷池运行数据列表，支持自定义筛选条件
// @Summary Get a paginated list of pooldatas by custom conditions
// @Description Returns a paginated list of pooldata based on query filters, including page number and size.
// @Tags pooldata
// @Accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListPooldatasReply{}
// @Router /api/v1/pooldata/list [post]
// @Security BearerAuth
func (h *pooldataHandler) List(c *gin.Context) {
	form := &types.ListPooldatasRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	pooldatas, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertPooldatas(pooldatas)
	if err != nil {
		response.Error(c, ecode.ErrListPooldata)
		return
	}

	response.Success(c, gin.H{
		"pooldatas": data,
		"total":     total,
	})
}

// getPooldataIDFromPath 从请求路径中解析借贷池运行数据 ID，返回 (id字符串, id数值, 是否应中止请求)
func getPooldataIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

// convertPooldata 将模型层的 Pooldata 结构体转换为响应层结构体
func convertPooldata(pooldata *model.Pooldata) (*types.PooldataObjDetail, error) {
	data := &types.PooldataObjDetail{}
	err := copier.Copy(data, pooldata)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	return data, nil
}

// convertPooldatas 将模型层的 Pooldata 列表批量转换为响应层结构体
func convertPooldatas(fromValues []*model.Pooldata) ([]*types.PooldataObjDetail, error) {
	toValues := []*types.PooldataObjDetail{}
	for _, v := range fromValues {
		data, err := convertPooldata(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}
