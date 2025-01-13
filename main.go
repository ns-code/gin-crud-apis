package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ns-code/gin-crud-apis/docs"
	"github.com/ns-code/gin-crud-apis/handlers"
	"github.com/ns-code/gin-crud-apis/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupDBConn() {
	// init models.USERDB and models.USERDBERR global variables
	err := models.ConnectUserDatabase()
	handlers.CheckErr(err)
}

func SetupRouter(r *gin.Engine) *gin.Engine {

	r.Use(cors.Default())

	SetupSwaggerDocs()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api")
	{
		v1.GET("users", handlers.GetUsers)
		v1.POST("users", handlers.AddUser)
		v1.PUT("users/:user_id", handlers.UpdateUser)
		v1.DELETE("users/:user_id", handlers.DeleteUser)
	}
	return r
}

func SetupSwaggerDocs() {
	docs.SwaggerInfo.Title = "gin-crud-apis"
	docs.SwaggerInfo.Description = "gin crud apis"
	docs.SwaggerInfo.Version = ""
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}
}

func main() {
	SetupDBConn()
	SetupRouter(gin.Default()).Run()
}
