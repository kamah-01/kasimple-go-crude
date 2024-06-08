// handlers/handlers.go
package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "simple-crud/database"
    "simple-crud/models"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)

    collection := database.GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    result, err := collection.InsertOne(ctx, user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(result)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    var user models.User
    collection := database.GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    collection := database.GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var user models.User
        cursor.Decode(&user)
        users = append(users, user)
    }

    json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    var user models.User
    _ = json.NewDecoder(r.Body).Decode(&user)

    collection := database.GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    collection := database.GetCollection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(result)
}
