package docker

import (
	"context"
	"io"

	// สำหรับตัวอย่างอ่าน stream
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type dockerService struct {
	cli *client.Client
}

var _ DockerDomain = (*dockerService)(nil)

func NewDockerService(cli *client.Client) DockerDomain {
	return &dockerService{cli: cli}
}

// Implementations (ตัวอย่างบางส่วน)

func (s *dockerService) GetInfo(ctx context.Context) (system.Info, error) {
	return s.cli.Info(ctx)
}

func (s *dockerService) GetVersion(ctx context.Context) (types.Version, error) {
	return s.cli.ServerVersion(ctx)
}

func (s *dockerService) ListContainers(ctx context.Context, options container.ListOptions) ([]container.Summary, error) {
	return s.cli.ContainerList(ctx, options)
}

func (s *dockerService) ListImages(ctx context.Context, options image.ListOptions) ([]image.Summary, error) {
	return s.cli.ImageList(ctx, options)
}

func (s *dockerService) PullImage(ctx context.Context, ref string, options image.PullOptions) (io.ReadCloser, error) {
	// ใน Service อาจจะอ่าน Stream และ Log เอง หรือส่งกลับให้ Handler
	// ตัวอย่างนี้ส่ง Stream กลับให้ Handler
	return s.cli.ImagePull(ctx, ref, options)
}
func (s *dockerService) BuildImage(ctx context.Context, options types.ImageBuildOptions) (io.ReadCloser, error) {
	response, err := s.cli.ImageBuild(ctx, options.Context, options)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

func (s *dockerService) PushImage(ctx context.Context, ref string, options image.PushOptions) (io.ReadCloser, error) {
	return s.cli.ImagePush(ctx, ref, options)
}

func (s *dockerService) CreateImage(ctx context.Context, options image.CreateOptions) (io.ReadCloser, error) {
	// The first parameter is parentReference (usually the image name like "alpine:latest")
	// We can pass an empty string or a default image if options doesn't specify
	return s.cli.ImageCreate(ctx, "", options)
}

func (s *dockerService) CreateContainer(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, containerName string) (container.CreateResponse, error) {
	return s.cli.ContainerCreate(ctx, config, hostConfig, nil, nil, containerName)
}

func (s *dockerService) StartContainer(ctx context.Context, containerID string, options container.StartOptions) error {
	return s.cli.ContainerStart(ctx, containerID, options)
}

func (s *dockerService) StopContainer(ctx context.Context, containerID string, timeout *int) error {
	options := container.StopOptions{}
	if timeout != nil {
		options.Timeout = timeout
	}
	return s.cli.ContainerStop(ctx, containerID, options)
}

func (s *dockerService) RemoveContainer(ctx context.Context, containerID string, options container.RemoveOptions) error {
	return s.cli.ContainerRemove(ctx, containerID, options)
}

func (s *dockerService) ExecContainerCreate(ctx context.Context, containerID string, config container.ExecOptions) (container.ExecCreateResponse, error) {
	return s.cli.ContainerExecCreate(ctx, containerID, config)
}

func (s *dockerService) ExecContainerStart(ctx context.Context, execID string, config container.ExecStartOptions) (types.HijackedResponse, error) {
	return s.cli.ContainerExecAttach(ctx, execID, config)
}

func (s *dockerService) ExecContainerInspect(ctx context.Context, execID string) (container.ExecInspect, error) {
	return s.cli.ContainerExecInspect(ctx, execID)
}

func (s *dockerService) ContainerLogs(ctx context.Context, containerID string, options container.LogsOptions) (io.ReadCloser, error) {
	return s.cli.ContainerLogs(ctx, containerID, options)
}

func (s *dockerService) CreateNetwork(ctx context.Context, name string, options network.CreateOptions) (network.CreateResponse, error) {
	return s.cli.NetworkCreate(ctx, name, options)
}

func (s *dockerService) CreateVolume(ctx context.Context, options volume.CreateOptions) (volume.Volume, error) {
	return s.cli.VolumeCreate(ctx, options)
}

func (s *dockerService) RemoveNetwork(ctx context.Context, networkID string) error {
	return s.cli.NetworkRemove(ctx, networkID)
}

func (s *dockerService) RemoveVolume(ctx context.Context, volumeID string) error {
	return s.cli.VolumeRemove(ctx, volumeID, false)
}

// ... Implement Method อื่นๆ
