package http

import (
	_ "clean-architecture-example/internal/delivery/http/docs" // Импорт для документации
	"clean-architecture-example/internal/domain"
	"clean-architecture-example/internal/usecase"
	"encoding/json"
	"net/http"
)

// @title			User API
// @version		1.0
// @description	This is a sample user API.
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	support@example.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/
type UserHandler struct {
	addUserUseCase   *usecase.AddUserUseCase
	listUsersUseCase *usecase.ListUsersUseCase
}

func NewUserHandler(addUserUseCase *usecase.AddUserUseCase, listUsersUseCase *usecase.ListUsersUseCase) *UserHandler {
	return &UserHandler{
		addUserUseCase:   addUserUseCase,
		listUsersUseCase: listUsersUseCase,
	}
}

// @Summary		Add a new user
// @Description	Add a new user to the system
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			user	body		domain.User	true	"User object"
// @Success		201		{string}	string		"Created"
// @Failure		400		{string}	string		"Bad Request"
// @Failure		500		{string}	string		"Internal Server Error"
// @Router			/users [post]
func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.addUserUseCase.Execute(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// @Summary		List all users
// @Description	Get a list of all users
// @Tags			users
// @Produce		json
// @Success		200	{array}		domain.User
// @Failure		500	{string}	string	"Internal Server Error"
// @Router			/users [get]
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.listUsersUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
