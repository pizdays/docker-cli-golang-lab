package docker

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/api/types/volume"
)

type DockerDomain interface {
	GetInfo(ctx context.Context) (system.Info, error)
	GetVersion(ctx context.Context) (types.Version, error)
	ListContainers(ctx context.Context, options container.ListOptions) ([]types.Container, error)
	ListImages(ctx context.Context, options image.ListOptions) ([]image.Summary, error)

	BuildImage(ctx context.Context, options types.ImageBuildOptions) (io.ReadCloser, error)
	PushImage(ctx context.Context, ref string, options image.PushOptions) (io.ReadCloser, error)
	CreateImage(ctx context.Context, options image.CreateOptions) (io.ReadCloser, error)

	CreateContainer(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, containerName string) (container.CreateResponse, error)
	StartContainer(ctx context.Context, containerID string, options container.StartOptions) error
	StopContainer(ctx context.Context, containerID string, timeout *int) error
	RemoveContainer(ctx context.Context, containerID string, options container.RemoveOptions) error

	PullImage(ctx context.Context, ref string, options image.PullOptions) (io.ReadCloser, error)

	ExecContainerCreate(ctx context.Context, containerID string, config container.ExecOptions) (types.IDResponse, error)
	ExecContainerStart(ctx context.Context, execID string, config container.ExecStartOptions) (types.HijackedResponse, error)
	ExecContainerInspect(ctx context.Context, execID string) (container.ExecInspect, error)
	ContainerLogs(ctx context.Context, containerID string, options container.LogsOptions) (io.ReadCloser, error)

	CreateNetwork(ctx context.Context, name string, options network.CreateOptions) (network.CreateResponse, error)
	RemoveNetwork(ctx context.Context, networkID string) error

	CreateVolume(ctx context.Context, options volume.CreateOptions) (volume.Volume, error)
	RemoveVolume(ctx context.Context, volumeID string) error

	// เพิ่ม Method อื่นๆ
}