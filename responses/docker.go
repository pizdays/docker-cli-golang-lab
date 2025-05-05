package responses

type ErrorResponse struct {
	Message string `json:"message"`
}


type InfoResponse struct {
	OSType        string `json:"os_type"`
	ServerVersion string `json:"server_version"`
	Containers    int64  `json:"containers"`
	Images        int64  `json:"images"`
}

type VersionResponse struct {
	Version    string `json:"version"`
	APIVersion string `json:"api_version"`
	GoVersion  string `json:"go_version"`
}

type ContainerResponse struct {
	ID     string   `json:"id"`
	Image  string   `json:"image"`
	Status string   `json:"status"`
	Names  []string `json:"names"`
}

type CreateContainerResponse struct {
	ID       string   `json:"id"`
	Warnings []string `json:"warnings,omitempty"`
}

type ExecResponse struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	ExitCode int  `json:"exit_code"`
}

// เพิ่ม Response Structs อื่นๆ เช่น ImageResponse, NetworkResponse, VolumeResponse