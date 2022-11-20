package main

import (
	"latihan1/controllers"
	"latihan1/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Hactiv8 Swagger
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:4444
// @BasePath  /api/v1
func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", controllers.Ping)
	}
	todos := v1.Group("/todos")
	{
		todos.POST("", controllers.CreateTodo)
		todos.GET("", controllers.GetAll)
		todos.GET("/:id", controllers.GetByID)
		todos.PUT("/:id", controllers.UpdateByID)
		todos.DELETE("/:id", controllers.DeleteByID)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":4444")
}
