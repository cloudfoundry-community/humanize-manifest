package main

// Manifest ...
type Manifest struct {
	Name           string              `yaml:"name,omitempty"` // Ref: https://bosh.io/docs/manifest-v2/#deployment
	DirectorUUID   string              `yaml:"director_uuid,omitempty"`
	Tags           interface{}         `yaml:"tags,omitempty"` // Ref: https://bosh.io/docs/manifest-v2/#tags
	InstanceGroups []InstanceGroup     `yaml:"instance_groups,omitempty"`
	Jobs           []InstanceGroup     `yaml:"jobs,omitempty"`     // deprecated v1 alias for 'instance_groups'
	Features       interface{}         `yaml:"features,omitempty"` // Ref: https://bosh.io/docs/manifest-v2/#features
	Update         Update              `yaml:"update,omitempty"`
	Addons         []Addon             `yaml:"addons,omitempty"`
	Properties     interface{}         `yaml:"properties,omitempty"`     // v1. Deprecated in favor of job properties
	AZs            []AZDefinition      `yaml:"azs,omitempty"`            // v2. only in Cloud-Config
	Networks       []NetworkDefinition `yaml:"networks,omitempty"`       // v1. Obsoleted by Cloud-Config, but useful for 'create-env'. Also in v2 in Cloud-Config
	ResourcePools  []VMProfile         `yaml:"resource_pools,omitempty"` // v1. Obsoleted by Cloud-Config VM Types, but useful for 'create-env'
	VMTypes        []VMType            `yaml:"vm_types,omitempty"`       // v2. only in Cloud-Config
	VMExtensions   []VMExtension       `yaml:"vm_extensions,omitempty"`  // v2. only in Cloud-Config
	DiskPools      []DiskProfile       `yaml:"disk_pools,omitempty"`     // v1. Obsoleted by Cloud-Config Disk Types, but useful for 'create-env'
	DiskTypes      []DiskProfile       `yaml:"disk_types,omitempty"`     // v2. only in Cloud-Config
	Compilation    CompilationConfig   `yaml:"compilation,omitempty"`    // v1. Obsoleted by Cloud-Config, but useful for 'create-env'. Also in v2 in Cloud-Config
	CloudProvider  CPIConfig           `yaml:"cloud_provider,omitempty"` // Only valid with 'create-env' (a.k.a. bosh-init)
	Variables      []Variable          `yaml:"variables,omitempty"`
	Releases       []Release           `yaml:"releases,omitempty"`
	Stemcells      []Stemcell          `yaml:"stemcells,omitempty"`
	// Cloud Config
}

// InstanceGroup ...
// Ref: https://bosh.io/docs/manifest-v2/#instance-groups
type InstanceGroup struct {
	Name               string         `yaml:"name,omitempty"`
	MigratedFrom       []MigratedFrom `yaml:"migrated_from,omitempty"`
	Instances          int            `yaml:"instances,omitempty"`
	Lifecycle          string         `yaml:"lifecycle,omitempty"`
	AZs                []string       `yaml:"azs,omitempty,flow"`
	Jobs               []Job          `yaml:"jobs,omitempty"`
	Templates          []Job          `yaml:"templates,omitempty"`  // deprecated v1 alias for 'jobs'
	Properties         interface{}    `yaml:"properties,omitempty"` // v1. Deprecated in favor of job properties
	Stemcell           string         `yaml:"stemcell,omitempty"`
	VMType             string         `yaml:"vm_type,omitempty"`
	ResourcePool       string         `yaml:"resource_pool,omitempty"` // v1 concept similar to 'vm_type'
	VMExtensions       []interface{}  `yaml:"vm_extensions,omitempty"`
	VMResources        []interface{}  `yaml:"vm_resources,omitempty"`
	PersistentDisk     int            `yaml:"persistent_disk,omitempty"`
	PersistentDiskType string         `yaml:"persistent_disk_type,omitempty"`
	PersistentDiskPool string         `yaml:"persistent_disk_pool,omitempty"` // v1 concept similar to 'persistent_disk_type'
	Env                interface{}    `yaml:"env,omitempty"`
	Networks           []Network      `yaml:"networks,omitempty"`
	Update             Update         `yaml:"update,omitempty"`
}

// MigratedFrom ...
// Ref: https://bosh.io/docs/migrated-from/#schema
type MigratedFrom struct {
	Name string `yaml:"name,omitempty"`
	AZ   string `yaml:"azs,omitempty"`
}

// Job ...
type Job struct {
	Name       string      `yaml:"name,omitempty"`
	Release    string      `yaml:"release,omitempty"`
	Provides   interface{} `yaml:"provides,omitempty"`
	Consumes   interface{} `yaml:"consumes,omitempty"`
	Properties interface{} `yaml:"properties,omitempty"`
}

// Network ...
type Network struct {
	Name      string   `yaml:"name,omitempty"`
	Default   []string `yaml:"default,omitempty,flow"`
	StaticIPs []string `yaml:"static_ips,omitempty"`
}

// Update ...
// Ref: https://bosh.io/docs/manifest-v2/#update
type Update struct {
	Serial          bool   `yaml:"serial,omitempty"`
	Canaries        int    `yaml:"canaries,omitempty"`
	CanaryWatchTime string `yaml:"canary_watch_time,omitempty"`
	MaxInFlight     int    `yaml:"max_in_flight,omitempty"`
	UpdateWatchTime string `yaml:"update_watch_time,omitempty"`
}

// Addon ...
// Ref: https://bosh.io/docs/manifest-v2/#addons
type Addon struct {
	Name    string      `yaml:"name,omitempty"`
	Jobs    []Job       `yaml:"jobs,omitempty"`
	Include interface{} `yaml:"include,omitempty"`
	Exclude interface{} `yaml:"exclude,omitempty"`
}

// AZDefinition ...
type AZDefinition struct {
	Name            string      `yaml:"name,omitempty"`
	CloudProperties interface{} `yaml:"cloud_properties,omitempty"`
}

// NetworkDefinition ...
// Ref: https://bosh.io/docs/deployment-manifest/#networks
type NetworkDefinition struct {
	Name            string             `yaml:"name,omitempty"` // 'manual', 'dynamic', or 'vip'
	Type            string             `yaml:"type,omitempty"`
	DNS             []string           `yaml:"dns,omitempty,flow"` // for 'dynamic' networks with no subnets only
	Subnets         []SubnetDefinition `yaml:"subnets,omitempty"`  // for 'manual' or 'dynamic' networks only
	CloudProperties interface{}        `yaml:"cloud_properties,omitempty"`
}

// SubnetDefinition ...
type SubnetDefinition struct {
	Range           string      `yaml:"range,omitempty"`   // for 'manual' networks only
	Gateway         string      `yaml:"gateway,omitempty"` // for 'manual' networks only
	DNS             []string    `yaml:"dns,omitempty,flow"`
	Reserved        []string    `yaml:"reserved,omitempty"` // for 'manual' networks only
	Static          []string    `yaml:"static,omitempty"`   // for 'manual' networks only
	AZ              string      `yaml:"az,omitempty"`
	AZs             []string    `yaml:"azs,omitempty,flow"`
	CloudProperties interface{} `yaml:"cloud_properties,omitempty"`
}

// VMProfile ...
// Ref: https://bosh.io/docs/deployment-manifest/#resource-pools
type VMProfile struct {
	Name            string                `yaml:"name,omitempty"`
	Network         string                `yaml:"network,omitempty"`
	Size            int                   `yaml:"size,omitempty"`
	Stemcell        BoshCreateEnvStemcell `yaml:"stemcell,omitempty"`
	CloudProperties interface{}           `yaml:"cloud_properties,omitempty"`
	Env             interface{}           `yaml:"env,omitempty"`
}

// VMType ...
// Ref: https://bosh.io/docs/cloud-config/#vm-types
type VMType struct {
	Name            string      `yaml:"name,omitempty"`
	CloudProperties interface{} `yaml:"cloud_properties,omitempty"`
}

// VMExtension ...
// Ref: https://bosh.io/docs/cloud-config/#vm-extensions
type VMExtension struct {
	Name            string      `yaml:"name,omitempty"`
	CloudProperties interface{} `yaml:"cloud_properties,omitempty"`
}

// BoshCreateEnvStemcell ...
type BoshCreateEnvStemcell struct {
	URL  string `yaml:"url,omitempty"`
	SHA1 string `yaml:"sha1,omitempty"`
}

// DiskProfile ...
// Ref: https://bosh.io/docs/deployment-manifest/#disk-pools
type DiskProfile struct {
	Name            string      `yaml:"name,omitempty"`
	DiskSize        int         `yaml:"disk_size,omitempty"`
	CloudProperties interface{} `yaml:"cloud_properties,omitempty"`
}

// CompilationConfig ...
// Ref. v1: https://bosh.io/docs/deployment-manifest/#compilation
type CompilationConfig struct {
	Workers             int         `yaml:"workers,omitempty"`
	AZ                  string      `yaml:"az,omitempty"`
	VMType              string      `yaml:"vm_type,omitempty"`
	VMResources         interface{} `yaml:"vm_resources,omitempty"`
	Network             string      `yaml:"network,omitempty"`
	ReuseCompilationVMs bool        `yaml:"reuse_compilation_vms,omitempty"`
	CloudProperties     interface{} `yaml:"cloud_properties,omitempty"`
	Env                 interface{} `yaml:"env,omitempty"`
}

// CPIConfig ...
type CPIConfig struct {
	Template   interface{} `yaml:"template,omitempty"`
	MBus       string      `yaml:"mbus,omitempty"`
	Cert       string      `yaml:"cert,omitempty"`
	Properties interface{} `yaml:"properties,omitempty"`
}

// Variable ...
// Ref: https://bosh.io/docs/manifest-v2/#variables
type Variable struct {
	Name    string          `yaml:"name,omitempty"`
	Type    string          `yaml:"type,omitempty"`
	Options VariableOptions `yaml:"options,omitempty"`
}

// VariableOptions ...
type VariableOptions struct {
	IsCA             bool     `yaml:"is_ca,omitempty"`
	CA               string   `yaml:"ca,omitempty"`
	CommonName       string   `yaml:"common_name,omitempty"`
	AlternativeNames []string `yaml:"alternative_names,omitempty"`
	ExtendedKeyUsage []string `yaml:"extended_key_usage,omitempty,flow"`
}

// Release ...
//
// Ref. v2: https://bosh.io/docs/manifest-v2/#releases
//
// Ref. v1: https://bosh.io/docs/deployment-manifest/#releases
// Ref. v1: https://bosh.io/docs/deployment-manifest/#bosh-init-stemcells
type Release struct {
	Name    string `yaml:"name,omitempty"`
	Version string `yaml:"version,omitempty"`
	URL     string `yaml:"url,omitempty"`
	SHA1    string `yaml:"sha1,omitempty"`
}

// Stemcell ...
// Ref. v2: https://bosh.io/docs/manifest-v2/#stemcells
type Stemcell struct {
	Alias   string `yaml:"alias,omitempty"`
	Name    string `yaml:"name,omitempty"` // 'alias' is preferred over 'name'
	Os      string `yaml:"os,omitempty"`
	Version string `yaml:"version,omitempty"`
}
