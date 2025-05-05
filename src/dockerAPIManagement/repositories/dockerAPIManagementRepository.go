package repositories

import (
	"github.com/docker-cli-golang-lab/src/dockerAPIManagement/handlers"
	"github.com/docker-cli-golang-lab/src/dockerAPIManagement/usecases"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	conn     *gorm.DB
}

// init Repository Handler
func NewRepositoryHandler(conn *gorm.DB) *handlers.Handler {
	useCase := usecases.NewUseCase(&Repository{conn})
	handler := handlers.NewHandler(useCase)
	return handler
}
