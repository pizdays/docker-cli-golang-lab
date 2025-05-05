package repositories

import (
	"github.com/docker-cli-golang-lab/services/docker"
	"github.com/docker-cli-golang-lab/src/dockerAPIManagement/handlers"
	"github.com/docker-cli-golang-lab/src/dockerAPIManagement/usecases"
	"github.com/docker/docker/client"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	conn *gorm.DB
}

// init Repository Handler
func NewRepositoryHandler(conn *gorm.DB, cli *client.Client) *handlers.Handler {
	useCase := usecases.NewUseCase(&Repository{conn}, docker.NewDockerService(cli))
	handler := handlers.NewHandler(useCase)
	return handler
}

// // Repository interface implementations
// func (r *Repository) GetServiceByID(id string) (interface{}, error) {
// 	// Implement fetching service from database
// 	return nil, nil
// }

// func (r *Repository) CreateService(service interface{}) error {
// 	// Implement service creation in database
// 	return nil
// }

// func (r *Repository) UpdateService(id string, service interface{}) error {
// 	// Implement service update in database
// 	return nil
// }

// func (r *Repository) DeleteService(id string) error {
// 	// Implement service deletion from database
// 	return nil
// }
