package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"backend/config"
	"backend/controller"
	"backend/repository"
	"backend/routes"
	"backend/service"
	"github.com/gin-gonic/gin"
)

// @title 学生信息管理系统 API
// @version 1.0
// @description 学生信息管理系统的 API 文档

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// 初始化数据库
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取当前工作目录失败: %v", err)
	}

	// 构建日志文件路径
	logDir := filepath.Join(dir, "log")
	// 确保日志目录存在
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("创建日志目录失败: %v", err)
	}
	logPath := filepath.Join(logDir, "app.log")
	fmt.Printf("日志文件路径: %s\n", logPath)

	// 初始化仓库
	studentRepo := repository.StudentRepository{DB: db}
	userRepo := repository.UserRepository{DB: db}

	// 初始化服务
	studentService := service.StudentService{StudentRepo: &studentRepo}
	authService := service.AuthService{UserRepo: &userRepo}

	// 初始化控制器
	studentController := controller.StudentController{StudentService: &studentService}
	authController := controller.AuthController{AuthService: &authService}

	// 创建 Gin 引擎
	router := gin.Default()

	// 挂载 CORS 中间件
	router.Use(CORSMiddleware())

	// 设置路由
	router = routes.SetupRouter(&authController, &studentController, &authService)

	// 打印所有注册的路由，用于调试
	for _, route := range router.Routes() {
		log.Printf("已注册路由 - 方法: %s，路径: %s，处理器: %s", route.Method, route.Path, route.Handler)
	}

	// 健康检查接口
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 获取端口号
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("服务将启动在端口 %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}

// CORSMiddleware CORS 中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") //用的时候改成前端服务器地址吧
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "false") //这个改成true，不然用不了

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
