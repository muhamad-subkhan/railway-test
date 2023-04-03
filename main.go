package main

import (
	"fmt"
	"net/http"
	"os"
	"party/database"
	"party/pkg/mysql"
	"party/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	mysql.Database()
	
	database.Migration()
	
	r := mux.NewRouter()

	routes.Routes(r.PathPrefix("/api").Subrouter())
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	var Port = os.Getenv("PORT")
	fmt.Println("Server Running on Port:"+ Port)

	http.ListenAndServe(":" + Port, r)
}