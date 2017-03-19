package driver

import (
	"crypto/tls"
	"net"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
)

// ServeTCP makes the handler to listen for request in a given TCP address.
// It also writes the spec file on the right directory for docker to read.
func (h Handler) ServeTCP(pluginName, addr string, tlsConfig *tls.Config) error {
	l, spec, err := newTCPListener(addr, pluginName, tlsConfig)
	if err != nil {
		return err
	}
	if spec != "" {
		defer os.Remove(spec)
	}
	return h.Serve(l)
}

// ServeUnix makes the handler to listen for requests in a unix socket.
// It also creates the socket file on the right directory for docker to read.
func (h Handler) ServeUnix(addr string, gid int) error {
	l, spec, err := newUnixListener(addr, gid)
	if err != nil {
		return err
	}
	if spec != "" {
		defer os.Remove(spec)
	}
	return h.Serve(l)
}

// Serve sets up the handler to serve requests on the passed in listener
func (h Handler) Serve(l net.Listener) error {
	server := http.Server{
		Addr:    l.Addr().String(),
		Handler: h.mux,
	}
	logrus.Info("Start tcp serve")
	return server.Serve(l)
}
