package driver

import "github.com/itoryio/docme/storage"

//DocmeMultihostDriver is driver implementation
type DocmeMultihostDriver struct {
	Storage *storage.Storage
}

func (docme *DocmeMultihostDriver) Create(r Request) Response {
	return Response{}
}

func (docme *DocmeMultihostDriver) List(Request) Response {
	return Response{}
}
func (docme *DocmeMultihostDriver) Get(Request) Response {
	return Response{}
}
func (docme *DocmeMultihostDriver) Remove(Request) Response {
	return Response{}
}
func (docme *DocmeMultihostDriver) Path(Request) Response {
	return Response{}
}
func (docme *DocmeMultihostDriver) Mount(MountRequest) Response {
	return Response{}
}
func (docme *DocmeMultihostDriver) Unmount(UnmountRequest) Response {
	return Response{}
}
func (docme *DocmeMultihostDriver) Capabilities(Request) Response {
	return Response{}
}
