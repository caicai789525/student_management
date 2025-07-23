package routes

import (
	"backend/controller"
	"backend/middleware"
	"backend/service"
	"github.com/gin-gonic/gin"
	"log"
)

func SetupRouter(authController *controller.AuthController, studentController *controller.StudentController, authService *service.AuthService) *gin.Engine {
	r := gin.Default()

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		log.Println("正在注册 /auth/login POST 路由")
		authGroup.POST("/login", authController.Login)
	}

	studentGroup := r.Group("/students")
	studentGroup.Use(middleware.AuthMiddleware(authService))
	{
		studentGroup.POST("create", studentController.CreateStudent)
		studentGroup.POST("update", studentController.UpdateStudent)
		studentGroup.POST("delete", studentController.DeleteStudent)
		studentGroup.GET("", studentController.GetStudents)
		// 修改为使用学号
		studentGroup.GET("/:studentID", studentController.GetStudent)
		studentGroup.PUT("/:studentID", studentController.UpdateStudent)
		studentGroup.DELETE("/:studentID", studentController.DeleteStudent)
	}

	// 打印所有注册的路由
	for _, route := range r.Routes() {
		log.Printf("已注册路由 - 方法: %s，路径: %s，处理器: %s", route.Method, route.Path, route.Handler)
	}

	return r
}
