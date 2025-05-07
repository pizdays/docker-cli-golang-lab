package models

// SystemInfo represents system information from Docker
type SystemInfo struct {
	ID                 string             `json:"ID"`
	Containers         int                `json:"Containers"`
	ContainersRunning  int                `json:"ContainersRunning"`
	ContainersPaused   int                `json:"ContainersPaused"`
	ContainersStopped  int                `json:"ContainersStopped"`
	Images             int                `json:"Images"`
	Driver             string             `json:"Driver"`
	DriverStatus       [][]string         `json:"DriverStatus"`
	SystemStatus       [][]string         `json:"SystemStatus"`
	Plugins            Plugins            `json:"Plugins"`
	MemoryLimit        bool               `json:"MemoryLimit"`
	SwapLimit          bool               `json:"SwapLimit"`
	KernelMemory       bool               `json:"KernelMemory"`
	CPUCfsPeriod       bool               `json:"CpuCfsPeriod"`
	CPUCfsQuota        bool               `json:"CpuCfsQuota"`
	CPUShares          bool               `json:"CPUShares"`
	CPUSet             bool               `json:"CPUSet"`
	PidsLimit          bool               `json:"PidsLimit"`
	IPv4Forwarding     bool               `json:"IPv4Forwarding"`
	BridgeNfIptables   bool               `json:"BridgeNfIptables"`
	BridgeNfIP6tables  bool               `json:"BridgeNfIp6tables"`
	Debug              bool               `json:"Debug"`
	NFd                int                `json:"NFd"`
	NGoroutines        int                `json:"NGoroutines"`
	SystemTime         string             `json:"SystemTime"`
	LoggingDriver      string             `json:"LoggingDriver"`
	CgroupDriver       string             `json:"CgroupDriver"`
	NEventsListener    int                `json:"NEventsListener"`
	KernelVersion      string             `json:"KernelVersion"`
	OperatingSystem    string             `json:"OperatingSystem"`
	OSVersion          string             `json:"OSVersion"`
	OSType             string             `json:"OSType"`
	Architecture       string             `json:"Architecture"`
	NCPU               int                `json:"NCPU"`
	MemTotal           int64              `json:"MemTotal"`
	IndexServerAddress string             `json:"IndexServerAddress"`
	RegistryConfig     RegistryConfig     `json:"RegistryConfig"`
	NMem               int                `json:"NMem"`
	GenericResources   []GenericResource  `json:"GenericResources"`
	DockerRootDir      string             `json:"DockerRootDir"`
	HTTPProxy          string             `json:"HttpProxy"`
	HTTPSProxy         string             `json:"HttpsProxy"`
	NoProxy            string             `json:"NoProxy"`
	Name               string             `json:"Name"`
	Labels             []string           `json:"Labels"`
	Experimental       bool               `json:"Experimental"`
	ServerVersion      string             `json:"ServerVersion"`
	ClusterStore       string             `json:"ClusterStore"`
	ClusterAdvertise   string             `json:"ClusterAdvertise"`
	Runtimes           map[string]Runtime `json:"Runtimes"`
	DefaultRuntime     string             `json:"DefaultRuntime"`
	Swarm              SwarmInfo          `json:"Swarm"`
	LiveRestoreEnabled bool               `json:"LiveRestoreEnabled"`
	Isolation          string             `json:"Isolation"`
	InitBinary         string             `json:"InitBinary"`
	ContainerdCommit   Commit             `json:"ContainerdCommit"`
	RuncCommit         Commit             `json:"RuncCommit"`
	InitCommit         Commit             `json:"InitCommit"`
	SecurityOptions    []string           `json:"SecurityOptions"`
	ProductLicense     string             `json:"ProductLicense"`
	Warnings           []string           `json:"Warnings"`
}

// Plugins represents Docker plugins
type Plugins struct {
	Volume  []string `json:"Volume"`
	Network []string `json:"Network"`
	Log     []string `json:"Log"`
}

// RegistryConfig represents Docker registry configuration
type RegistryConfig struct {
	InsecureRegistryCIDRs []string               `json:"InsecureRegistryCIDRs"`
	IndexConfigs          map[string]IndexConfig `json:"IndexConfigs"`
	Mirrors               []string               `json:"Mirrors"`
}

// IndexConfig represents Docker index configuration
type IndexConfig struct {
	Name     string   `json:"Name"`
	Mirrors  []string `json:"Mirrors"`
	Secure   bool     `json:"Secure"`
	Official bool     `json:"Official"`
}

// GenericResource represents a generic resource
type GenericResource struct {
	NamedResourceSpec    NamedResourceSpec    `json:"NamedResourceSpec"`
	DiscreteResourceSpec DiscreteResourceSpec `json:"DiscreteResourceSpec"`
}

// NamedResourceSpec represents a named resource specification
type NamedResourceSpec struct {
	Kind  string `json:"Kind"`
	Value string `json:"Value"`
}

// DiscreteResourceSpec represents a discrete resource specification
type DiscreteResourceSpec struct {
	Kind  string `json:"Kind"`
	Value int64  `json:"Value"`
}

// Runtime represents a container runtime
type Runtime struct {
	Path string   `json:"path"`
	Args []string `json:"runtimeArgs"`
}

// SwarmInfo represents Docker Swarm information
type SwarmInfo struct {
	NodeID           string  `json:"NodeID"`
	NodeAddr         string  `json:"NodeAddr"`
	LocalNodeState   string  `json:"LocalNodeState"`
	ControlAvailable bool    `json:"ControlAvailable"`
	Error            string  `json:"Error"`
	RemoteManagers   []Peer  `json:"RemoteManagers"`
	Nodes            int     `json:"Nodes"`
	Managers         int     `json:"Managers"`
	Cluster          Cluster `json:"Cluster"`
}

// Peer represents a Swarm peer
type Peer struct {
	NodeID string `json:"NodeID"`
	Addr   string `json:"Addr"`
}

// Cluster represents a Swarm cluster
type Cluster struct {
	ID                     string       `json:"ID"`
	Version                SwarmVersion `json:"Version"`
	CreatedAt              string       `json:"CreatedAt"`
	UpdatedAt              string       `json:"UpdatedAt"`
	Spec                   Spec         `json:"Spec"`
	TLSInfo                TLSInfo      `json:"TLSInfo"`
	RootRotationInProgress bool         `json:"RootRotationInProgress"`
}

// SwarmVersion represents a version
type SwarmVersion struct {
	Index uint64 `json:"Index"`
}

// Spec represents a specification
type Spec struct {
	Name          string            `json:"Name"`
	Labels        map[string]string `json:"Labels"`
	Orchestration Orchestration     `json:"Orchestration"`
	Raft          Raft              `json:"Raft"`
	Dispatcher    Dispatcher        `json:"Dispatcher"`
	CAConfig      CAConfig          `json:"CAConfig"`
	TaskDefaults  TaskDefaults      `json:"TaskDefaults"`
}

// Orchestration represents orchestration settings
type Orchestration struct {
	TaskHistoryRetentionLimit int `json:"TaskHistoryRetentionLimit"`
}

// Raft represents Raft settings
type Raft struct {
	SnapshotInterval           uint64 `json:"SnapshotInterval"`
	KeepOldSnapshots           uint64 `json:"KeepOldSnapshots"`
	LogEntriesForSlowFollowers uint64 `json:"LogEntriesForSlowFollowers"`
	ElectionTick               int    `json:"ElectionTick"`
	HeartbeatTick              int    `json:"HeartbeatTick"`
}

// Dispatcher represents dispatcher settings
type Dispatcher struct {
	HeartbeatPeriod int64 `json:"HeartbeatPeriod"`
}

// CAConfig represents CA configuration
type CAConfig struct {
	NodeCertExpiry int64        `json:"NodeCertExpiry"`
	ExternalCAs    []ExternalCA `json:"ExternalCAs"`
}

// ExternalCA represents external CA configuration
type ExternalCA struct {
	Protocol string            `json:"Protocol"`
	URL      string            `json:"URL"`
	Options  map[string]string `json:"Options"`
	CACert   string            `json:"CACert"`
}

// TaskDefaults represents task default settings
type TaskDefaults struct {
	LogDriver LogDriver `json:"LogDriver"`
}

// LogDriver represents log driver settings
type LogDriver struct {
	Name    string            `json:"Name"`
	Options map[string]string `json:"Options"`
}

// TLSInfo represents TLS information
type TLSInfo struct {
	TrustRoot           string `json:"TrustRoot"`
	CertIssuerSubject   string `json:"CertIssuerSubject"`
	CertIssuerPublicKey string `json:"CertIssuerPublicKey"`
}

// Commit represents a commit
type Commit struct {
	ID       string `json:"ID"`
	Expected string `json:"Expected"`
}
