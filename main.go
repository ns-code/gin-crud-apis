package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"github.com/ns-code/gin-crud-apis/docs"
	"github.com/ns-code/gin-crud-apis/handlers"
	"github.com/ns-code/gin-crud-apis/models"
	"github.com/ns-code/gin-crud-apis/util"

	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func SetupDBConn() {
	// init models.USERDB and models.USERDBERR global variables
	err := models.ConnectUserDatabase()
	util.CheckErr(err, "users.db start error")
}

func RunMuxServer() {

	muxRouter := mux.NewRouter()
	SetupSwaggerDocs()
	muxRouter.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	// muxRouter := muxr.PathPrefix("/api").Subrouter()	
	muxRouter.HandleFunc("/api/users", handlers.GetUsers).Methods(http.MethodGet)
	muxRouter.HandleFunc("/api/users", handlers.AddUser).Methods(http.MethodPost)
	muxRouter.HandleFunc("/api/users/{userId}", handlers.UpdateUser).Methods(http.MethodPut)
	muxRouter.HandleFunc("/api/users/{userId}", handlers.DeleteUser).Methods(http.MethodDelete)

	fmt.Println(">> Mux Server is starting at port 8080")

	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:4200"},
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowCredentials: true,
    })

    handler := c.Handler(muxRouter)
    log.Fatal(http.ListenAndServe(":8080", handler))
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
	RunMuxServer()
}
