package main

import (
	"github.com/gin-contrib/cors"
	"github.com/ns-code/gin-crud-apis/docs"
	"github.com/gin-gonic/gin"
	"github.com/ns-code/gin-crud-apis/handlers"
	"github.com/ns-code/gin-crud-apis/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(r *gin.Engine) *gin.Engine {

	r.Use(cors.Default())

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api")
	{

		v1.GET("users", handlers.GetUsers)
		v1.POST("users", handlers.AddUser)
		v1.PUT("users/:user_id", handlers.UpdateUser)
		v1.DELETE("users/:user_id", handlers.DeleteUser)
	}
	return r;
}

func main() {

	// @Summary get users
	// @Description get string by ID
	docs.SwaggerInfo.Title = "gin-crud-apis"
	docs.SwaggerInfo.Description = "gin crud apis"
	docs.SwaggerInfo.Version = ""
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}


	err := models.ConnectUserDatabase()
	handlers.CheckErr(err)

	r := gin.Default()
	SetupRouter(r).Run()
}
