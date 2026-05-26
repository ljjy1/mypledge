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

var _ ContractHandler = (*contractHandler)(nil)

// ContractHandler defining the handler interface
type ContractHandler interface {
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
	UpdateByID(c *gin.Context)
	GetByID(c *gin.Context)
	List(c *gin.Context)
}

type contractHandler struct {
	iDao dao.ContractDao
}

// NewContractHandler creating the handler interface
func NewContractHandler() ContractHandler {
	return &contractHandler{
		iDao: dao.NewContractDao(
			database.GetDB(), // db driver is mysql
			cache.NewContractCache(database.GetCacheType()),
		),
	}
}

// Create a new contract
// @Summary Create a new contract
// @Description Creates a new contract entity using the provided data in the request body.
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.CreateContractRequest true "contract information"
// @Success 200 {object} types.CreateContractReply{}
// @Router /api/v1/contract [post]
// @Security BearerAuth
func (h *contractHandler) Create(c *gin.Context) {
	form := &types.CreateContractRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	contract := &model.Contract{}
	err = copier.Copy(contract, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateContract)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, contract)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": contract.ID})
}

// DeleteByID delete a contract by id
// @Summary Delete a contract by id
// @Description Deletes a existing contract identified by the given id in the path.
// @Tags contract
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeleteContractByIDReply{}
// @Router /api/v1/contract/{id} [delete]
// @Security BearerAuth
func (h *contractHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getContractIDFromPath(c)
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

// UpdateByID update a contract by id
// @Summary Update a contract by id
// @Description Updates the specified contract by given id in the path, support partial update.
// @Tags contract
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateContractByIDRequest true "contract information"
// @Success 200 {object} types.UpdateContractByIDReply{}
// @Router /api/v1/contract/{id} [put]
// @Security BearerAuth
func (h *contractHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getContractIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateContractByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	contract := &model.Contract{}
	err = copier.Copy(contract, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDContract)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, contract)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID get a contract by id
// @Summary Get a contract by id
// @Description Gets detailed information of a contract specified by the given id in the path.
// @Tags contract
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetContractByIDReply{}
// @Router /api/v1/contract/{id} [get]
// @Security BearerAuth
func (h *contractHandler) GetByID(c *gin.Context) {
	_, id, isAbort := getContractIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	contract, err := h.iDao.GetByID(ctx, id)
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

	data := &types.ContractObjDetail{}
	err = copier.Copy(data, contract)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDContract)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	response.Success(c, gin.H{"contract": data})
}

// List get a paginated list of contracts by custom conditions
// @Summary Get a paginated list of contracts by custom conditions
// @Description Returns a paginated list of contract based on query filters, including page number and size.
// @Tags contract
// @Accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListContractsReply{}
// @Router /api/v1/contract/list [post]
// @Security BearerAuth
func (h *contractHandler) List(c *gin.Context) {
	form := &types.ListContractsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	contracts, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertContracts(contracts)
	if err != nil {
		response.Error(c, ecode.ErrListContract)
		return
	}

	response.Success(c, gin.H{
		"contracts": data,
		"total":     total,
	})
}

func getContractIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

func convertContract(contract *model.Contract) (*types.ContractObjDetail, error) {
	data := &types.ContractObjDetail{}
	err := copier.Copy(data, contract)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	return data, nil
}

func convertContracts(fromValues []*model.Contract) ([]*types.ContractObjDetail, error) {
	toValues := []*types.ContractObjDetail{}
	for _, v := range fromValues {
		data, err := convertContract(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}
