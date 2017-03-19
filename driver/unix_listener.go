package driver

import (
	"net"
	"os"
	"path/filepath"

	"github.com/docker/go-connections/sockets"
)

//create new unix.sock and return listener
func newUnixListener(pluginName string, gid int) (net.Listener, string, error) {
	path, err := fullSocketAddress(pluginName)
	if err != nil {
		return nil, "", err
	}
	listener, err := sockets.NewUnixSocket(path, gid)
	if err != nil {
		return nil, "", err
	}
	return listener, path, nil
}

//create all dir file under Docker /run/docker/plugin and return .sock path
func fullSocketAddress(address string) (string, error) {
	if err := os.MkdirAll(pluginSockDir, 0755); err != nil {
		return "", err
	}
	if filepath.IsAbs(address) {
		return address, nil
	}
	return filepath.Join(pluginSockDir, address+".sock"), nil
}
