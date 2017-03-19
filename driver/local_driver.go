package driver

import (
	"github.com/itoryio/docme/storage"

	"path/filepath"

	"github.com/Sirupsen/logrus"
)

//DocmeLocalDriver is driver implementation
type DocmeLocalDriver struct {
	Storage *storage.Storage
}

//Create method
func (docme *DocmeLocalDriver) Create(r Request) Response {
	rec := storage.Record{
		Name:        r.Name,
		Options:     r.Options,
		Mountpoint:  filepath.Join(DefaultDockerRootDirectory, r.Name),
		Connections: 0,
	}
	docme.Storage.Create(rec)
	logrus.Infoln(rec)
	logrus.Infoln("Docker call CREATE cmd")
	logrus.Infoln(r.Name)
	logrus.Infoln(r.Options)
	return Response{}
}

//List method
func (docme *DocmeLocalDriver) List(Request) Response {
	return Response{}
}

//Get method
func (docme *DocmeLocalDriver) Get(Request) Response {
	logrus.Infoln("Docker call GET cmd")
	return Response{}
}

//Remove method
func (docme *DocmeLocalDriver) Remove(Request) Response {
	return Response{}
}

//Path method
func (docme *DocmeLocalDriver) Path(Request) Response {
	return Response{}
}

//Mount method
func (docme *DocmeLocalDriver) Mount(MountRequest) Response {
	return Response{}
}

//Unmount method
func (docme *DocmeLocalDriver) Unmount(UnmountRequest) Response {
	return Response{}
}

//Capabilities method
func (docme *DocmeLocalDriver) Capabilities(Request) Response {
	return Response{}
}
