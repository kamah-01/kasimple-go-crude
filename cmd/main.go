package main

import (
    "log"
    "net/http"
    "simple-crud/database"
    "simple-crud/routes"
    "github.com/gorilla/mux"
)

func main() {
    database.Connect()

    r := mux.NewRouter()
    routes.RegisterRoutes(r)

    log.Println("Server started on: http://localhost:8080")
    http.ListenAndServe(":8080", r)
}
