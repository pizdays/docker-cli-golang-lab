package handlers

import (
	"github.com/docker-cli-golang-lab/src/dockerAPIManagement/domains"
	"github.com/gin-gonic/gin"
)

// init handler
type Handler struct {
	UseCase domains.UseCase
}

// init handler
func NewHandler(UseCase domains.UseCase) *Handler {
	return &Handler{
		UseCase: UseCase,
	}
}

func (h *Handler) CreateService(c *gin.Context) {

	return
}

func (h *Handler) UpdateService(c *gin.Context) {

	return
}

func (h *Handler) DeleteService(c *gin.Context) {

	return
}
