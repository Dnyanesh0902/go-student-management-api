package routes

import (
	"student-management-api/controllers"
	"student-management-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api")

	authRoutes := api.Group("/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}
	studentRoutes := api.Group("/students")
	studentRoutes.Use(middleware.AuthMiddleware())
	{
		studentRoutes.POST("/",middleware.AuthorizeRole("Admin"), controllers.CreateStudent)
		studentRoutes.GET("/", controllers.GetStudents)
		studentRoutes.GET("/:id", controllers.GetStudentByID)
		studentRoutes.PUT("/:id", controllers.UpdateStudent)
		studentRoutes.DELETE("/:id",middleware.AuthorizeRole("Admin"), controllers.DeleteStudent)
	}
}