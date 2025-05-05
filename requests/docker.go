package requests

// สำหรับ Port Mapping

type CreateContainerRequest struct {
	ImageName string   `json:"image_name" binding:"required"`
	ContainerName string `json:"container_name"`
	Cmd []string `json:"cmd"`
	Env []string `json:"env"`
	PortMappings []string `json:"port_mappings"` // e.g., ["8080:80", "443:443"]
}

type PullImageRequest struct {
	ImageName string `json:"image_name" binding:"required"`
}

type ExecRequest struct {
	Command []string `json:"command" binding:"required"` // e.g., ["ls", "-l", "/"]
}

// เพิ่ม Request Structs อื่นๆ เช่น CreateNetworkRequest, CreateVolumeRequest
/*
type CreateNetworkRequest struct {
    Name    string `json:"name" binding:"required"`
    Driver  string `json:"driver"`
}
*/