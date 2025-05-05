package domains

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/system"
)

// business logic
type UseCase interface {
	// Docker Service Operations
	GetInfo(ctx context.Context) (system.Info, error)
	GetVersion(ctx context.Context) (types.Version, error)
	ListContainers(ctx context.Context, options container.ListOptions) ([]types.Container, error)
	ListImages(ctx context.Context, options image.ListOptions) ([]image.Summary, error)
	
	// CreateContainer(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, containerName string) (container.CreateResponse, error)
	// StartContainer(ctx context.Context, containerID string, options container.StartOptions) error
	// StopContainer(ctx context.Context, containerID string, timeout *int) error
	// RemoveContainer(ctx context.Context, containerID string, options container.RemoveOptions) error
	
	// PullImage(ctx context.Context, ref string, options image.PullOptions) (io.ReadCloser, error)
	// BuildImage(ctx context.Context, options types.ImageBuildOptions) (io.ReadCloser, error)
	// PushImage(ctx context.Context, ref string, options image.PushOptions) (io.ReadCloser, error)
	// CreateImage(ctx context.Context, options image.CreateOptions) (io.ReadCloser, error)
	
	// ExecContainerCreate(ctx context.Context, containerID string, config container.ExecOptions) (types.IDResponse, error)
	// ExecContainerStart(ctx context.Context, execID string, config container.ExecStartOptions) (types.HijackedResponse, error)
	// ExecContainerInspect(ctx context.Context, execID string) (container.ExecInspect, error)
	// ContainerLogs(ctx context.Context, containerID string, options container.LogsOptions) (io.ReadCloser, error)
	
	// CreateNetwork(ctx context.Context, name string, options network.CreateOptions) (network.CreateResponse, error)
	// RemoveNetwork(ctx context.Context, networkID string) error
	
	// CreateVolume(ctx context.Context, options volume.CreateOptions) (volume.Volume, error)
	// RemoveVolume(ctx context.Context, volumeID string) error

	// Service operations for your application
	// GetService()
	// GetServiceByID()
	// CreateService()
	// UpdateService()
	// DeleteService()

}

// Repository interface for database operations
type Repository interface {
	// Add any database related methods here
		// GetServiceByID(id string) (interface{}, error)
		// CreateService(service interface{}) error
		// UpdateService(id string, service interface{}) error
		// DeleteService(id string) error
}
