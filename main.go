package main

import (
	"fmt"
	"log"
	"net/http"

	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ns-code/gin-crud-apis/docs"
	"github.com/ns-code/gin-crud-apis/handlers"
	"github.com/ns-code/gin-crud-apis/models"
	"github.com/ns-code/gin-crud-apis/util"
)

func SetupDBConn() {
	// init models.USERDB and models.USERDBERR global variables
	err := models.ConnectUserDatabase()
	util.CheckErr(err, "users.db start error")
}

func RunMuxServer() {

	r := mux.NewRouter()
	
	SetupSwaggerDocs()

	r.HandleFunc("/api/users", handlers.GetUsers).Methods("GET")
	// r.HandleFunc("/api/users", handlers.AddUser).Methods("POST")
	// r.HandleFunc("/api/users/:user_id", handlers.UpdateUser).Methods("PUT")
	// r.HandleFunc("/api/users/:user_id", handlers.DeleteUser).Methods("DELETE")

/*     corsOptions := gorillahandlers.CORSOptions{
        AllowedOrigins: []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
        AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Allowed request headers
    }
	
	handler := gorillahandlers.CORS(corsOptions)(r) */

	fmt.Println(">> Mux Server is starting at port 8080")
	log.Fatal(http.ListenAndServe(":8080", gorillahandlers.CORS()(r)))
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
