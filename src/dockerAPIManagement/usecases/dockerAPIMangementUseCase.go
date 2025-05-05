package usecases

import (
	"context"
	"fmt"
	"log"

	serviceDocker "github.com/docker-cli-golang-lab/services/docker"
	"github.com/docker-cli-golang-lab/src/dockerAPIManagement/domains"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/system"
)

type UseCase struct {
	Repo          domains.Repository
	DockerService serviceDocker.DockerDomain
}

// init usecase
func NewUseCase(repo domains.Repository, dockerService serviceDocker.DockerDomain) domains.UseCase {
	return &UseCase{
		Repo:          repo,
		DockerService: dockerService,
	}
}
// Docker Service Operations
func (u *UseCase) GetInfo(ctx context.Context) (system.Info, error) {
	// Logic: แค่เรียก Infrastructure Service ตรงๆ
	info, err := u.DockerService.GetInfo(ctx)
	if err != nil {
		// Usecase อาจจะ log error หรือแปลง error ให้มีความหมายทาง Business มากขึ้นก่อน return
		log.Printf("Usecase Error GetDockerInfo: %v", err)
		return system.Info{}, fmt.Errorf("failed to get docker info: %w", err) // Wrap error หรือ return error ใหม่
	}
	return info, nil
}

func (u *UseCase) GetVersion(ctx context.Context) (types.Version, error) {
	version, err := u.DockerService.GetVersion(ctx)
	if err != nil {
		log.Printf("Usecase Error GetDockerVersion: %v", err)
		return types.Version{}, fmt.Errorf("failed to get docker version: %w", err)
	}
	return version, nil
}

func (u *UseCase) ListContainers(ctx context.Context, options container.ListOptions) ([]types.Container, error) {
	containers, err := u.DockerService.ListContainers(ctx, options)
	if err != nil {
		log.Printf("Usecase Error ListContainers: %v", err)
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}
	return containers, nil
}

// func (u *UseCase) ListImages(ctx context.Context, options image.ListOptions) ([]image.Summary, error) {
	
	
	
// 	images, err := u.DockerService.ListImages(ctx, options)
// 	if err != nil {
// 		log.Printf("Usecase Error ListImages: %v", err)
// 		return nil, fmt.Errorf("failed to list images: %w", err)
// 	}
// 	return images, nil
// }



// Service operations
// func (u *UseCase) CreateService() {
// Implement your service creation logic here
// 	return
// }

// func (u *UseCase) UpdateService() {
// Implement your service update logic here
// 	return
// }

// func (u *UseCase) DeleteService() {
	// Implement service deletion logic here
// 	return
// }

// func (u *UseCase) DeployStack() {
// 	// Implement your stack deployment logic here
// 	return
// }
