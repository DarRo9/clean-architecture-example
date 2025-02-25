package main

import (
	"database/sql"
	"log"
	"net/http"

	myhttp "clean-architecture-example/internal/delivery/http"
	"clean-architecture-example/internal/repository"
	"clean-architecture-example/internal/usecase"

	_ "clean-architecture-example/internal/delivery/http/docs" // Import for documentation

	_ "github.com/mattn/go-sqlite3"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title User API
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	// Initialize the database
	db := initDB()
	defer db.Close()

	// Initialize repository, use cases, and handler
	userHandler := setupUserHandler(db)

	// Set up routes
	setupRoutes(userHandler)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	return db
}

func setupUserHandler(db *sql.DB) *myhttp.UserHandler {
	userRepo := repository.NewSQLiteUserRepository(db)
	userRepo.CreateTable()
	addUserUseCase := usecase.NewAddUserUseCase(userRepo)
	listUsersUseCase := usecase.NewListUsersUseCase(userRepo)
	return myhttp.NewUserHandler(addUserUseCase, listUsersUseCase)
}

func setupRoutes(userHandler *myhttp.UserHandler) {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userHandler.AddUser(w, r)
		case http.MethodGet:
			userHandler.ListUsers(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Swagger documentation endpoint
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
}
