package controller

import (
	"backend/model"
	"backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SuccessResponse 成功响应结构体
type SuccessResponse struct {
	Message string `json:"message"`
}

// ErrorResponse 错误响应结构体
type ErrorResponse struct {
	Error string `json:"error"`
}

// TokenResponse Token 响应结构体
type TokenResponse struct {
	Token string `json:"token"`
}

// LoginRequest 用户登录请求结构体
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthController struct {
	AuthService *service.AuthService
}

// Register 用户注册
// @Summary 用户注册
// @Description 新用户注册接口
// @Tags 认证
// @Accept json
// @Produce json
// @Param user body model.User true "用户信息"
// @Success 201 {object} SuccessResponse "注册成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 500 {object} ErrorResponse "服务器内部错误"
// @Router /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.AuthService.Register(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录接口，返回 JWT Token
// @Tags 认证
// @Accept json
// @Produce json
// @Param loginData body LoginRequest true "登录信息"
// @Success 200 {object} TokenResponse "登录成功，返回 Token"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 401 {object} ErrorResponse "认证失败"
// @Router /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var loginData LoginRequest
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	token, err := c.AuthService.Login(loginData.Username, loginData.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, TokenResponse{Token: token})
}
