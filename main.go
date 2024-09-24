package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("HomeHandler Worlks!")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    //r.HandleFunc("/products", ProductsHandler)
    //r.HandleFunc("/articles", ArticlesHandler)
    http.Handle("/", r)


    connStr := "user=admin password=root dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err == nil{
		fmt.Printf("Error db connections")
		return;
	}

    http.ListenAndServe(":8080", r)
}