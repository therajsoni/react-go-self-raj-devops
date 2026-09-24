package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"office-user-management/database"
	"office-user-management/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	_, err = database.UserCollection.InsertOne(ctx, user)

	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	cursor, err := database.UserCollection.Find(ctx, bson.M{})

	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(ctx)

	var users []models.User

	err = cursor.All(ctx, &users)

	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	if users == nil {
		users = []models.User{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	err = database.UserCollection.FindOne(
		ctx,
		bson.M{"_id": objectID},
	).Decode(&user)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	update := bson.M{
		"name":          user.Name,
		"username":      user.Username,
		"email":         user.Email,
		"phoneNo":       user.PhoneNo,
		"joiningDate":   user.JoiningDate,
		"dob":           user.DOB,
		"officeTimeIn":  user.OfficeTimeIn,
		"officeTimeOut": user.OfficeTimeOut,
		"photoBase64":   user.PhotoBase64,
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	result, err := database.UserCollection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": update},
	)

	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	user.ID = objectID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	result, err := database.UserCollection.DeleteOne(
		ctx,
		bson.M{"_id": objectID},
	)

	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "User deleted successfully",
		},
	)
}