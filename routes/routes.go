// routes/routes.go
package routes

import (
    "simple-crud/handlers"
    "github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
    r.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
    r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
}
