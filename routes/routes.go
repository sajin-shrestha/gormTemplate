package routes

import (
	"github.com/gorilla/mux"
	"github.com/sajin-shrestha/gormTemplate/controllers"
	"github.com/sajin-shrestha/gormTemplate/middlewares"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()

	// Apply CORS middleware globally
	r.Use(middlewares.CORS)

	// Define routes without JWT authentication
	r.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	// r.HandleFunc("/api/login", controllers.Login).Methods("POST")
	r.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", controllers.GetUserByID).Methods("GET")
	r.HandleFunc("/api/users/{id}", controllers.DeleteUserByID).Methods("DELETE")

	// Create a subrouter for routes that require JWT authentication
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTAuth)
	// api.HandleFunc("/users/{id}", controllers.DeleteUserByID).Methods("DELETE")

	return r

	// r := mux.NewRouter()

	// // Apply CORS middleware globally
	// r.Use(middlewares.CORS)

	// // Create a subrouter for API routes
	// api := r.PathPrefix("/api").Subrouter()

	// // Apply JWTAuth middleware to the API subrouter
	// api.Use(middlewares.JWTAuth)

	// // Define routes within the API subrouter
	// api.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	// api.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	// api.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	// api.HandleFunc("/users/{id}", controllers.DeleteUserByID).Methods("DELETE")

	// return r
}
