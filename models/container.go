package models

// ContainerSummary represents a container summary from Docker
type ContainerSummary struct {
	ID              string            `json:"Id"`
	Names           []string          `json:"Names"`
	Image           string            `json:"Image"`
	ImageID         string            `json:"ImageID"`
	Command         string            `json:"Command"`
	Created         int64             `json:"Created"`
	State           string            `json:"State"`
	Status          string            `json:"Status"`
	Ports           []Port            `json:"Ports"`
	Labels          map[string]string `json:"Labels"`
	NetworkSettings NetworkSettings   `json:"NetworkSettings"`
	Mounts          []Mount           `json:"Mounts"`
}

// Port represents a port mapping
type Port struct {
	IP          string `json:"IP"`
	PrivatePort uint16 `json:"PrivatePort"`
	PublicPort  uint16 `json:"PublicPort"`
	Type        string `json:"Type"`
}

// NetworkSettings represents the network settings of a container
type NetworkSettings struct {
	Networks map[string]Network `json:"Networks"`
}

// Network represents a network configuration
type Network struct {
	NetworkID string `json:"NetworkID"`
	IPAddress string `json:"IPAddress"`
	Gateway   string `json:"Gateway"`
}

// Mount represents a mount point
type Mount struct {
	Type        string `json:"Type"`
	Name        string `json:"Name"`
	Source      string `json:"Source"`
	Destination string `json:"Destination"`
	Driver      string `json:"Driver"`
	Mode        string `json:"Mode"`
	RW          bool   `json:"RW"`
}
