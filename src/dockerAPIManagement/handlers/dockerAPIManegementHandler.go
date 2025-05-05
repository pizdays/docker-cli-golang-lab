package handlers

import (
	"net/http"

	"github.com/docker-cli-golang-lab/src/dockerAPIManagement/domains"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
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
func (h *Handler) GetInfo(c *gin.Context) {
	info, err := h.UseCase.GetInfo(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"info": info,
	})
}

func (h *Handler) GetVersion(c *gin.Context) {
	version, err := h.UseCase.GetVersion(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"version": version,
	})
 }

func (h *Handler) ListContainers(c *gin.Context) {
	containers, err := h.UseCase.ListContainers(c.Request.Context(), container.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"containers": containers,
	})
}

func (h *Handler) ListImages(c *gin.Context) {
	images, err := h.UseCase.ListImages(c.Request.Context(), image.ListOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"images": images,
	})
}



// func (h *Handler) CreateContainer(c *gin.Context) {
// 	var request struct {
// 		Config        *container.Config    `json:"config" binding:"required"`
// 		HostConfig    *container.HostConfig `json:"host_config"`
// 		ContainerName string               `json:"container_name"`
// 	}
	
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
	
// 	response, err := h.UseCase.CreateContainer(c.Request.Context(), request.Config, request.HostConfig, request.ContainerName)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
	
// 	c.JSON(http.StatusCreated, gin.H{
// 		"container": response,
// 	})
// }

// func (h *Handler) RemoveContainer(c *gin.Context) {
// 	containerID := c.Param("id")
// 	if containerID == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "container ID is required",
// 		})
// 		return
// 	}
	
// 	var options container.RemoveOptions
// 	if err := c.ShouldBindQuery(&options); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
	
// 	if err := h.UseCase.RemoveContainer(c.Request.Context(), containerID, options); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
	
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Container removed successfully",
// 	})
// }

// func (h *Handler) PullImage(c *gin.Context) {
// 	var request struct {
// 		Reference string            `json:"reference" binding:"required"`
// 		Options   image.PullOptions `json:"options"`
// 	}
	
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
	
// 	reader, err := h.UseCase.PullImage(c.Request.Context(), request.Reference, request.Options)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	defer reader.Close()
	
// 	c.Writer.Header().Set("Content-Type", "application/json")
// 	c.Status(http.StatusOK)
	
// 	// Stream the response directly to the client
// 	_, err = io.Copy(c.Writer, reader)
// 	if err != nil {
// 		// Can't send error response at this point as headers are already sent
// 		// Just log the error
// 		log.Printf("Error streaming pull image response: %v", err)
// 	}
// }

// func (h *Handler) BuildImage(c *gin.Context) {
// 	var options types.ImageBuildOptions
	
// 	if err := c.ShouldBindJSON(&options); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
	
// 	reader, err := h.UseCase.BuildImage(c.Request.Context(), options)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	defer reader.Close()
	
// 	c.Writer.Header().Set("Content-Type", "application/json")
// 	c.Status(http.StatusOK)
	
// 	// Stream the response directly to the client
// 	_, err = io.Copy(c.Writer, reader)
// 	if err != nil {
// 		log.Printf("Error streaming build image response: %v", err)
// 	}
// }

// func (h *Handler) PushImage(c *gin.Context) {
// 	var request struct {
// 		Reference string            `json:"reference" binding:"required"`
// 		Options   image.PushOptions `json:"options"`
// 	}
	
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
	
// 	reader, err := h.UseCase.PushImage(c.Request.Context(), request.Reference, request.Options)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	defer reader.Close()
	
// 	c.Writer.Header().Set("Content-Type", "application/json")
// 	c.Status(http.StatusOK)
	
// 	// Stream the response directly to the client
// 	_, err = io.Copy(c.Writer, reader)
// 	if err != nil {
// 		log.Printf("Error streaming push image response: %v", err)
// 	}
// }


// func (h *Handler) CreateService(c *gin.Context) {
// 	return
// }

// func (h *Handler) UpdateService(c *gin.Context) {
// 	return
// }

// func (h *Handler) DeleteService(c *gin.Context) {
// 	return
// }


