package driver

import "net/http"

const (
	// DefaultDockerRootDirectory is the default directory where volumes will be created.
	DefaultDockerRootDirectory = "/mnt/docme/volumes"
	//pluginSpecDir              = "/run/docker/plugins"
	pluginSpecDir = "/etc/docker/plugins"
	pluginSockDir = "/run/docker/plugins"

	manifest         = `{"Implements": ["VolumeDriver"]}`
	createPath       = "/VolumeDriver.Create"
	getPath          = "/VolumeDriver.Get"
	listPath         = "/VolumeDriver.List"
	removePath       = "/VolumeDriver.Remove"
	hostVirtualPath  = "/VolumeDriver.Path"
	mountPath        = "/VolumeDriver.Mount"
	unmountPath      = "/VolumeDriver.Unmount"
	capabilitiesPath = "/VolumeDriver.Capabilities"
)

// Request is the structure that docker's requests are deserialized to.
type Request struct {
	Name    string
	Options map[string]string `json:"Opts,omitempty"`
}

// Response is the strucutre that the plugin's responses are serialized to.
type Response struct {
	Mountpoint   string
	Err          string
	Volumes      []*Volume
	Volume       *Volume
	Capabilities Capability
}

// MountRequest structure for a volume mount request
type MountRequest struct {
	Name string
	ID   string
}

// UnmountRequest structure for a volume unmount request
type UnmountRequest struct {
	Name string
	ID   string
}

// Volume represents a volume object for use with `Get` and `List` requests
type Volume struct {
	Name       string
	Mountpoint string
	Status     map[string]interface{}
}

// Capability represents the list of capabilities a volume driver can return
type Capability struct {
	Scope string
}

// Driver represent the interface a driver must fulfill.
type Driver interface {
	Create(Request) Response
	List(Request) Response
	Get(Request) Response
	Remove(Request) Response
	Path(Request) Response
	Mount(MountRequest) Response
	Unmount(UnmountRequest) Response
	Capabilities(Request) Response
}

// Handler forwards requests and responses between the docker daemon and the plugin.
type Handler struct {
	driver Driver
	mux    *http.ServeMux
}

type actionHandler func(Request) Response
type mountActionHandler func(MountRequest) Response
type unmountActionHandler func(UnmountRequest) Response
