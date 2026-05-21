package routes


import (
	
	"student-management-api/controllers"
	"github.com/gin-gonic/gin"
)


func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	studentRoutes := api.Group("/students") 
	{
		studentRoutes.POST("/", controllers.CreateStudent)
		studentRoutes.GET("/", controllers.GetStudent)
		studentRoutes.GET("/:id", controllers.GetStudentById)
		studentRoutes.PUT("/:id", controllers.UpdateStudent)
		studentRoutes.DELETE("/:id", controllers.DeleteStudent)
	};

	userRoutes := api.Group("/users")
	{
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUserByID)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}
}


