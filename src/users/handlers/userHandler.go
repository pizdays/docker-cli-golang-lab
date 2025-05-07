package handlers

import (
	"net/http"
	"strconv"

	"github.com/docker-cli-golang-lab/models"
	"github.com/docker-cli-golang-lab/requests"
	"github.com/docker-cli-golang-lab/responses"
	"github.com/docker-cli-golang-lab/src/users/domains"
	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	userUseCase domains.UserUseCase
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userUseCase domains.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// GetUsers handles GET /api/users
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} responses.UserResponse
// @Failure 500 {object} map[string]string "error"
// @Router /api/v1/users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userUseCase.GetUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to response format
	userResponses := make([]responses.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = responses.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
			Active:   user.Active,
		}
	}

	c.JSON(http.StatusOK, responses.UsersResponse{
		Users: userResponses,
		Count: len(userResponses),
	})
}

// GetUserByID handles GET /api/users/:id
// @Summary Get a user by ID
// @Description Retrieve a user by their unique ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} map[string]string "error"
// @Failure 404 {object} map[string]string "error"
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.userUseCase.GetUserByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toUserResponse(*user))
}

// CreateUser handles POST /api/users
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body requests.UserRequest true "User details"
// @Success 201 {object} responses.UserResponse
// @Failure 400 {object} map[string]string "error"
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var userReq requests.UserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert request to model
	user := models.User{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	if err := h.userUseCase.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, toUserResponse(user))
}

// UpdateUser handles PUT /api/users/:id
// @Summary Update a user by ID
// @Description Update a user's details by their unique ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body requests.UserRequest true "User details"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} map[string]string "error"
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var userReq requests.UserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert request to model
	user := models.User{
		ID:       uint(id),
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	if err := h.userUseCase.UpdateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toUserResponse(user))
}

// DeleteUser handles DELETE /api/users/:id
// @Summary Delete a user by ID
// @Description Delete a user by their unique ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]string "message"
// @Failure 400 {object} map[string]string "error"
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	if err := h.userUseCase.DeleteUser(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

// Login handles POST /api/users/login
// @Summary Login to the system
// @Description Authenticate a user and generate a JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body requests.UserRequest true "Login credentials"
// @Success 200 {object} responses.UserLoginResponse
// @Failure 400 {object} map[string]string "error"
// @Failure 401 {object} map[string]string "error"
// @Router /api/v1/users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var credentials requests.UserRequest
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userUseCase.AuthenticateUser(c.Request.Context(), credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.UserLoginResponse{
		User:  toUserResponse(*user),
		Token: "mock-jwt-token", // In a real app, you would generate a JWT token here
	})
}

// Helper to convert a model to a response struct
func toUserResponse(user models.User) responses.UserResponse {
	return responses.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
