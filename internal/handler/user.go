package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"goal-advanced-layout/api"
	"goal-advanced-layout/internal/biz"
)

func NewUserHandler(handler *Handler, uu *biz.UserUsecase) *UserHandler {
	return &UserHandler{
		Handler: handler,
		uu:      uu,
	}
}

type UserHandler struct {
	*Handler
	uu *biz.UserUsecase
}

// Register godoc
// @Summary 用户注册
// @Schemes
// @Description 目前只支持邮箱登录
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body api.RegisterRequest true "params"
// @Success 200 {object} api.Response
// @Router /register [post]
func (h *UserHandler) Register(ctx *gin.Context) {
	req := new(api.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		api.Fail(ctx, http.StatusBadRequest, api.ErrBadRequest, nil)
		return
	}

	if err := h.uu.Register(ctx, req); err != nil {
		h.logger.WithContext(ctx).Error("userService.Register error", zap.Error(err))
		api.Fail(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	api.Succeed(ctx, nil)
}

// Login godoc
// @Summary 账号登录
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param request body api.LoginRequest true "params"
// @Success 200 {object} api.LoginResponse
// @Router /login [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	var req api.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.Fail(ctx, http.StatusBadRequest, api.ErrBadRequest, nil)
		return
	}

	token, err := h.uu.Login(ctx, &req)
	if err != nil {
		api.Fail(ctx, http.StatusUnauthorized, api.ErrUnauthorized, nil)
		return
	}
	api.Succeed(ctx, api.LoginResponseData{
		AccessToken: token,
	})
}

// GetProfile godoc
// @Summary 获取用户信息
// @Schemes
// @Description
// @Tags 用户模块
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} api.GetProfileResponse
// @Router /user [get]
func (h *UserHandler) GetProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		api.Fail(ctx, http.StatusUnauthorized, api.ErrUnauthorized, nil)
		return
	}

	user, err := h.uu.GetProfile(ctx, userId)
	if err != nil {
		api.Fail(ctx, http.StatusBadRequest, api.ErrBadRequest, nil)
		return
	}

	api.Succeed(ctx, user)
}

func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)

	var req api.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.Fail(ctx, http.StatusBadRequest, api.ErrBadRequest, nil)
		return
	}

	if err := h.uu.UpdateProfile(ctx, userId, &req); err != nil {
		api.Fail(ctx, http.StatusInternalServerError, api.ErrInternalServerError, nil)
		return
	}

	api.Succeed(ctx, nil)
}
