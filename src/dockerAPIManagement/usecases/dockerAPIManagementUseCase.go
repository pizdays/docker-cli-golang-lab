package usecases

import (
	"context"
	"fmt"
	"io"
	"log"

	serviceDocker "github.com/docker-cli-golang-lab/services/docker"
	"github.com/docker-cli-golang-lab/src/dockerAPIManagement/domains"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/api/types/volume"
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

func (u *UseCase) ListContainers(ctx context.Context, options container.ListOptions) ([]container.Summary, error) {
	containers, err := u.DockerService.ListContainers(ctx, options)
	if err != nil {
		log.Printf("Usecase Error ListContainers: %v", err)
		return nil, fmt.Errorf("failed to list containers: %w", err)
	}
	return containers, nil
}

func (u *UseCase) ListImages(ctx context.Context, options image.ListOptions) ([]image.Summary, error) {

	images, err := u.DockerService.ListImages(ctx, options)
	if err != nil {
		log.Printf("Usecase Error ListImages: %v", err)
		return nil, fmt.Errorf("failed to list images: %w", err)
	}
	return images, nil
}
func (u *UseCase) CreateContainer(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, containerName string) (container.CreateResponse, error) {
	response, err := u.DockerService.CreateContainer(ctx, config, hostConfig, containerName)
	if err != nil {
		log.Printf("Usecase Error CreateContainer: %v", err)
		return container.CreateResponse{}, fmt.Errorf("failed to create container: %w", err)
	}
	return response, nil
}

func (u *UseCase) StartContainer(ctx context.Context, containerID string, options container.StartOptions) error {
	err := u.DockerService.StartContainer(ctx, containerID, options)
	if err != nil {
		log.Printf("Usecase Error StartContainer: %v", err)
		return fmt.Errorf("failed to start container: %w", err)
	}
	return nil
}

func (u *UseCase) StopContainer(ctx context.Context, containerID string, timeout *int) error {
	err := u.DockerService.StopContainer(ctx, containerID, timeout)
	if err != nil {
		log.Printf("Usecase Error StopContainer: %v", err)
		return fmt.Errorf("failed to stop container: %w", err)
	}
	return nil
}

func (u *UseCase) RemoveContainer(ctx context.Context, containerID string, options container.RemoveOptions) error {
	err := u.DockerService.RemoveContainer(ctx, containerID, options)
	if err != nil {
		log.Printf("Usecase Error RemoveContainer: %v", err)
		return fmt.Errorf("failed to remove container: %w", err)
	}
	return nil
}

func (u *UseCase) PullImage(ctx context.Context, ref string, options image.PullOptions) (io.ReadCloser, error) {
	reader, err := u.DockerService.PullImage(ctx, ref, options)
	if err != nil {
		log.Printf("Usecase Error PullImage: %v", err)
		return nil, fmt.Errorf("failed to pull image: %w", err)
	}
	return reader, nil
}

func (u *UseCase) BuildImage(ctx context.Context, options types.ImageBuildOptions) (io.ReadCloser, error) {
	reader, err := u.DockerService.BuildImage(ctx, options)
	if err != nil {
		log.Printf("Usecase Error BuildImage: %v", err)
		return nil, fmt.Errorf("failed to build image: %w", err)
	}
	return reader, nil
}

func (u *UseCase) PushImage(ctx context.Context, ref string, options image.PushOptions) (io.ReadCloser, error) {
	reader, err := u.DockerService.PushImage(ctx, ref, options)
	if err != nil {
		log.Printf("Usecase Error PushImage: %v", err)
		return nil, fmt.Errorf("failed to push image: %w", err)
	}
	return reader, nil
}

func (u *UseCase) CreateImage(ctx context.Context, options image.CreateOptions) (io.ReadCloser, error) {
	reader, err := u.DockerService.CreateImage(ctx, options)
	if err != nil {
		log.Printf("Usecase Error CreateImage: %v", err)
		return nil, fmt.Errorf("failed to create image: %w", err)
	}
	return reader, nil
}

func (u *UseCase) ExecContainerCreate(ctx context.Context, containerID string, config container.ExecOptions) (container.ExecCreateResponse, error) {
	response, err := u.DockerService.ExecContainerCreate(ctx, containerID, config)
	if err != nil {
		log.Printf("Usecase Error ExecContainerCreate: %v", err)
		return container.ExecCreateResponse{}, fmt.Errorf("failed to create exec instance: %w", err)
	}
	return response, nil
}

func (u *UseCase) ExecContainerStart(ctx context.Context, execID string, config container.ExecStartOptions) (types.HijackedResponse, error) {
	response, err := u.DockerService.ExecContainerStart(ctx, execID, config)
	if err != nil {
		log.Printf("Usecase Error ExecContainerStart: %v", err)
		return types.HijackedResponse{}, fmt.Errorf("failed to start exec instance: %w", err)
	}
	return response, nil
}

func (u *UseCase) ExecContainerInspect(ctx context.Context, execID string) (container.ExecInspect, error) {
	inspect, err := u.DockerService.ExecContainerInspect(ctx, execID)
	if err != nil {
		log.Printf("Usecase Error ExecContainerInspect: %v", err)
		return container.ExecInspect{}, fmt.Errorf("failed to inspect exec instance: %w", err)
	}
	return inspect, nil
}

func (u *UseCase) ContainerLogs(ctx context.Context, containerID string, options container.LogsOptions) (io.ReadCloser, error) {
	reader, err := u.DockerService.ContainerLogs(ctx, containerID, options)
	if err != nil {
		log.Printf("Usecase Error ContainerLogs: %v", err)
		return nil, fmt.Errorf("failed to get container logs: %w", err)
	}
	return reader, nil
}

func (u *UseCase) CreateNetwork(ctx context.Context, name string, options network.CreateOptions) (network.CreateResponse, error) {
	response, err := u.DockerService.CreateNetwork(ctx, name, options)
	if err != nil {
		log.Printf("Usecase Error CreateNetwork: %v", err)
		return network.CreateResponse{}, fmt.Errorf("failed to create network: %w", err)
	}
	return response, nil
}

func (u *UseCase) RemoveNetwork(ctx context.Context, networkID string) error {
	err := u.DockerService.RemoveNetwork(ctx, networkID)
	if err != nil {
		log.Printf("Usecase Error RemoveNetwork: %v", err)
		return fmt.Errorf("failed to remove network: %w", err)
	}
	return nil
}

func (u *UseCase) CreateVolume(ctx context.Context, options volume.CreateOptions) (volume.Volume, error) {
	vol, err := u.DockerService.CreateVolume(ctx, options)
	if err != nil {
		log.Printf("Usecase Error CreateVolume: %v", err)
		return volume.Volume{}, fmt.Errorf("failed to create volume: %w", err)
	}
	return vol, nil
}

func (u *UseCase) RemoveVolume(ctx context.Context, volumeID string) error {
	err := u.DockerService.RemoveVolume(ctx, volumeID)
	if err != nil {
		log.Printf("Usecase Error RemoveVolume: %v", err)
		return fmt.Errorf("failed to remove volume: %w", err)
	}
	return nil
}

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
