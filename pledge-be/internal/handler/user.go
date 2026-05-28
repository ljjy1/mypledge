package handler

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-dev-frame/sponge/pkg/copier"
	"github.com/go-dev-frame/sponge/pkg/gin/middleware"
	"github.com/go-dev-frame/sponge/pkg/gin/response"
	"github.com/go-dev-frame/sponge/pkg/logger"
	"github.com/go-dev-frame/sponge/pkg/utils"

	"pledge-be/internal/auth"
	"pledge-be/internal/cache"
	"pledge-be/internal/dao"
	"pledge-be/internal/database"
	"pledge-be/internal/ecode"
	"pledge-be/internal/model"
	"pledge-be/internal/types"
)

var _ UserHandler = (*userHandler)(nil)

// UserHandler 用户处理器接口
type UserHandler interface {
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
	UpdateByID(c *gin.Context)
	GetByID(c *gin.Context)
	List(c *gin.Context)
	Login(c *gin.Context)
	Register(c *gin.Context)
}

// userHandler 用户数据处理实现
type userHandler struct {
	iDao dao.UserDao
}

// NewUserHandler 创建用户处理器实例
func NewUserHandler() UserHandler {
	return &userHandler{
		iDao: dao.NewUserDao(
			database.GetDB(), // db driver is mysql
			cache.NewUserCache(database.GetCacheType()),
		),
	}
}

// Create 创建新的用户记录
// @Summary Create a new user
// @Description Creates a new user entity using the provided data in the request body.
// @Tags user
// @Accept json
// @Produce json
// @Param data body types.CreateUserRequest true "user information"
// @Success 200 {object} types.CreateUserReply{}
// @Router /api/v1/user [post]
// @Security BearerAuth
func (h *userHandler) Create(c *gin.Context) {
	form := &types.CreateUserRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	user := &model.User{}
	err = copier.Copy(user, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateUser)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, user)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": user.ID})
}

// DeleteByID 根据 ID 删除用户记录
// @Summary Delete a user by id
// @Description Deletes a existing user identified by the given id in the path.
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeleteUserByIDReply{}
// @Router /api/v1/user/{id} [delete]
// @Security BearerAuth
func (h *userHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getUserIDFromPath(c)
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

// UpdateByID 根据 ID 更新用户记录
// @Summary Update a user by id
// @Description Updates the specified user by given id in the path, support partial update.
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateUserByIDRequest true "user information"
// @Success 200 {object} types.UpdateUserByIDReply{}
// @Router /api/v1/user/{id} [put]
// @Security BearerAuth
func (h *userHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getUserIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateUserByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	user := &model.User{}
	err = copier.Copy(user, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDUser)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, user)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID 根据 ID 查询用户记录
// @Summary Get a user by id
// @Description Gets detailed information of a user specified by the given id in the path.
// @Tags user
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetUserByIDReply{}
// @Router /api/v1/user/{id} [get]
// @Security BearerAuth
func (h *userHandler) GetByID(c *gin.Context) {
	_, id, isAbort := getUserIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	user, err := h.iDao.GetByID(ctx, id)
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

	data := &types.UserObjDetail{}
	err = copier.Copy(data, user)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDUser)
		return
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	response.Success(c, gin.H{"user": data})
}

// List 分页查询用户列表，支持自定义筛选条件
// @Summary Get a paginated list of users by custom conditions
// @Description Returns a paginated list of user based on query filters, including page number and size.
// @Tags user
// @Accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListUsersReply{}
// @Router /api/v1/user/list [post]
// @Security BearerAuth
func (h *userHandler) List(c *gin.Context) {
	form := &types.ListUsersRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	// 将前端习惯的 1-based 页码转为系统内部的 0-based 分页
	if form.Params.Page > 0 {
		form.Params.Page--
	}

	ctx := middleware.WrapCtx(c)
	users, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertUsers(users)
	if err != nil {
		response.Error(c, ecode.ErrListUser)
		return
	}

	response.Success(c, gin.H{
		"users": data,
		"total": total,
	})
}

// Login 用户登录认证，验证用户名密码并返回 JWT 令牌
// @Summary User login
// @Description Authenticates user and returns a JWT token.
// @Tags user
// @Accept json
// @Produce json
// @Param data body types.LoginRequest true "login credentials"
// @Success 200 {object} types.LoginReply{}
// @Router /api/v1/login [post]
func (h *userHandler) Login(c *gin.Context) {
	form := &types.LoginRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	user, err := h.iDao.GetByLogin(c, form.Login)
	if err != nil {
		response.Error(c, ecode.ErrLogin)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		response.Error(c, ecode.ErrLogin)
		return
	}

	token, err := auth.GenerateToken(strconv.FormatUint(user.ID, 10))
	if err != nil {
		response.Error(c, ecode.InternalServerError)
		return
	}

	response.Success(c, gin.H{"token": token})
}

// Register 用户注册，包含字段格式验证和唯一性检查
// @Summary User register
// @Description Registers a new user with validation rules.
// @Tags user
// @Accept json
// @Produce json
// @Param data body types.RegisterRequest true "registration info"
// @Success 200 {object} types.RegisterReply{}
// @Router /api/v1/register [post]
func (h *userHandler) Register(c *gin.Context) {
	form := &types.RegisterRequest{}
	if err := c.ShouldBindJSON(form); err != nil {
		response.Error(c, ecode.InvalidParams)
		return
	}

	// 验证登录账号：3-10位字母或数字
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{3,10}$", form.Login); !ok {
		response.Error(c, ecode.ErrRegisterLoginFormat)
		return
	}

	// 验证密码：8-20位
	if ok, _ := regexp.MatchString("^.{8,20}$", form.Password); !ok {
		response.Error(c, ecode.ErrRegisterPwdFormat)
		return
	}

	// 验证昵称：1-12位
	if ok, _ := regexp.MatchString("^.{1,12}$", form.Nike); !ok {
		response.Error(c, ecode.ErrRegisterNikeFormat)
		return
	}

	// 检查登录账号唯一性
	if _, err := h.iDao.GetByLogin(c, form.Login); err == nil {
		response.Error(c, ecode.ErrRegisterLoginExists)
		return
	}

	// bcrypt 加密密码
	hashed, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("bcrypt error", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	// 创建用户
	user := &model.User{
		Login:    form.Login,
		Nike:     form.Nike,
		Password: string(hashed),
	}
	if err := h.iDao.Create(c, user); err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": user.ID})
}

// getUserIDFromPath 从请求路径中解析用户 ID，返回 (id字符串, id数值, 是否应中止请求)
func getUserIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

// convertUser 将模型层的 User 结构体转换为响应层结构体
func convertUser(user *model.User) (*types.UserObjDetail, error) {
	data := &types.UserObjDetail{}
	err := copier.Copy(data, user)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	return data, nil
}

// convertUsers 将模型层的 User 列表批量转换为响应层结构体
func convertUsers(fromValues []*model.User) ([]*types.UserObjDetail, error) {
	toValues := []*types.UserObjDetail{}
	for _, v := range fromValues {
		data, err := convertUser(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}
