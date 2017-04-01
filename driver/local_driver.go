package driver

import (
	"os"

	"github.com/itoryio/docme/storage"

	"path/filepath"

	log "github.com/Sirupsen/logrus"
)

//DocmeLocalDriver is driver implementation
type DocmeLocalDriver struct {
	Storage *storage.Storage
}

//Create method
func (docme *DocmeLocalDriver) Create(r Request) Response {

	log.Debugln("Docker call CREATE cmd !!")
	log.Debugln(r.Name)
	log.Debugln(r.Options)

	rec := storage.Record{
		Name:        r.Name,
		Options:     r.Options,
		Mountpoint:  filepath.Join(DefaultDockerRootDirectory, r.Name),
		Connections: 0,
	}
	docme.Storage.Create(rec)

	return Response{
		Err: "",
	}
}

//List method
func (docme *DocmeLocalDriver) List(r Request) Response {

	log.Debugln("Docker call LIST cmd")
	recordset := docme.Storage.List()
	log.Debugln(recordset)
	volumes := make([]*Volume, 0)
	for _, item := range recordset {
		volumes = append(volumes, &Volume{
			Name:       item.Name,
			Mountpoint: item.Mountpoint,
		})
	}
	return Response{
		Err:     "",
		Volumes: volumes,
	}
}

//Get method
func (docme *DocmeLocalDriver) Get(r Request) Response {

	log.Debugln("Docker call GET cmd")
	log.Debugln(r)
	recordset, err := docme.Storage.Get(r.Name)
	if err != nil {
		return Response{
			Err: err.Error(),
		}
	}

	return Response{
		Volume: &Volume{
			Name:       recordset.Name,
			Mountpoint: recordset.Mountpoint,
			Status:     recordset.Status,
		},
		Err: "",
	}
}

//Remove method
func (docme *DocmeLocalDriver) Remove(r Request) Response {
	log.Debugln("Docker call Remove cmd")
	log.Debugln(r)

	recordset, err := docme.Storage.Get(r.Name)
	if err != nil {
		return Response{
			Err: err.Error(),
		}
	}

	if recordset.Connections > 0 {
		return Response{
			Err: "Volume still used",
		}

	}

	err = docme.Storage.Delete(r.Name)
	if err != nil {
		return Response{
			Err: err.Error(),
		}
	}

	return Response{
		Err: "",
	}
}

//Path method
func (docme *DocmeLocalDriver) Path(r Request) Response {
	log.Debugln("Docker call Path cmd")
	log.Debugln(r)

	recordset, err := docme.Storage.Get(r.Name)
	if err != nil {
		return Response{
			Err: err.Error(),
		}
	}

	return Response{
		Err:        "",
		Mountpoint: recordset.Mountpoint,
	}
}

//Mount method
func (docme *DocmeLocalDriver) Mount(r MountRequest) Response {
	log.Debugln("Docker call Mount cmd")
	log.Debugln(r)
	recordset, err := docme.Storage.Get(r.Name)
	if err != nil {
		return Response{
			Err: err.Error(),
		}
	}

	//First connection
	if recordset.Connections == 0 {
		log.Debug("Create new dir: " + recordset.Mountpoint)
		// 1. create a dir
		if _, err := os.Stat(recordset.Mountpoint); os.IsNotExist(err) {
			log.Debug("Dir [" + recordset.Mountpoint + "] is IsNotExist")
			err = os.Mkdir(recordset.Mountpoint, 0666)
			if err != nil {
				log.Error(err.Error())
				return Response{
					Err: err.Error(),
				}
			}
		}
		// 2. @TODO create a mount NFS
	}
	// increase connection counter
	recordset.Connections++

	docme.Storage.Update(recordset.Name, recordset)

	return Response{
		Err:        "",
		Mountpoint: recordset.Mountpoint,
	}
}

//Unmount method
func (docme *DocmeLocalDriver) Unmount(r UnmountRequest) Response {
	log.Debugln("Docker call Unmount cmd")
	log.Debugln(r)
	recordset, err := docme.Storage.Get(r.Name)
	if err != nil {
		return Response{
			Err: err.Error(),
		}
	}

	if recordset.Connections == 0 {
		log.Error("Something strange")
		return Response{
			Err: "Cannot unmount dir",
		}
	}

	// decrease connection counter
	recordset.Connections--

	docme.Storage.Update(recordset.Name, recordset)

	return Response{
		Err: "",
	}
}

//Capabilities method
func (docme *DocmeLocalDriver) Capabilities(r Request) Response {
	log.Debugln("Docker call Capabilities cmd")
	log.Debugln(r)
	return Response{}
}

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}
