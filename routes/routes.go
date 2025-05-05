package routes

import (
	"os"
	"strings"

	databases "github.com/docker-cli-golang-lab/databases"
	_ "github.com/docker-cli-golang-lab/docs"
	dockerAPIManagementRepo "github.com/docker-cli-golang-lab/src/dockerAPIManagement/repositories"
	"github.com/gin-contrib/gzip"

	// NewRepositoryHandler
	// ""gitlab.com/dol-api-service/src/volumeAadt/usecases

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	r.GET("/public/*filepath", func(c *gin.Context) {
		filepath := c.Param("filepath")
		if strings.HasSuffix(filepath, ".png") {
			c.Header("Content-Type", "image/png")
		} else if strings.HasSuffix(filepath, ".jpg") {
			c.Header("Content-Type", "image/jpg")
		} else if strings.HasSuffix(filepath, ".pdf") {
			c.Header("Content-Type", "application/pdf")
		}
		c.File("public" + filepath)
	})

	r.GET("/storages/*filepath", func(c *gin.Context) {
		filepath := c.Param("filepath")
		if strings.HasSuffix(filepath, ".csv") {
			c.Header("Content-Type", "text/csv")
		} else if strings.HasSuffix(filepath, ".xlsx") {
			c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		} else if strings.HasSuffix(filepath, ".pdf") {
			c.Header("Content-Type", "application/pdf")
		} else if strings.HasSuffix(filepath, ".html") {
			c.Header("Content-Type", "application/html")
		} else if strings.HasSuffix(filepath, ".zip") {
			c.Header("Content-Type", "application/zip")
		} else if strings.HasSuffix(filepath, ".png") {
			c.Header("Content-Type", "image/png")
		} else if strings.HasSuffix(filepath, ".jpg") {
			c.Header("Content-Type", "image/jpg")
		} else if strings.HasSuffix(filepath, ".docx") {
			c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
		}
		c.File("storages" + filepath)
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// middlewares.AuthorizeJWT()
	v1 := selectAPIPath(r, os.Getenv("ENV"))
	{

		dockerAPIManagement := v1.Group("/docker")
		{
			handler := dockerAPIManagementRepo.NewRepositoryHandler(databases.DB)

			dockerAPIManagement.POST("", handler.CreateService)

		}


	}
	return r
}

func selectAPIPath(r *gin.Engine, env string) *gin.RouterGroup {
	if env == "dev" {
		return r.Group("api/v1")
	}

	return r.Group("/v1")
}
