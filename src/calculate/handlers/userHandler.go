package handlers

import (
	"github.com/docker-cli-golang-lab/src/users/domains"
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

// // GetUsers handles GET /api/users
// // @Summary Get all users
// // @Description Retrieve a list of all users
// // @Tags users
// // @Accept json
// // @Produce json
// // @Success 200 {array} responses.UserResponse
// // @Failure 500 {object} map[string]string "error"
// // @Router /api/v1/users [get]
// func (h *UserHandler) CalculateArea(c *gin.Context) {

// 	c.JSON(http.StatusOK, responses.UsersResponse{
// 		Users: userResponses,
// 		Count: len(userResponses),
// 	})
// }
