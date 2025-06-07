package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/layan2k/go-bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql" 
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	port := "8000"

	http.Handle("/", r)
	log.Println("Server is running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
