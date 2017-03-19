package driver

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
)

const activatePath = "/Plugin.Activate"

// NewHandler initializes the request handler with a driver implementation.
func NewHandler(driver Driver) *Handler {

	mux := http.NewServeMux()

	mux.HandleFunc(activatePath, func(w http.ResponseWriter, r *http.Request) {
		logrus.Infoln("Docker call ACTIVATEPATH")
		w.Header().Set("Content-Type", DefaultContentTypeV1_1)
		fmt.Fprintln(w, manifest)
	})

	h := &Handler{driver, mux}
	h.initMux()
	return h
}

//Map our driver function to route by specification
func (h *Handler) initMux() {
	h.handle(createPath, func(req Request) Response {
		return h.driver.Create(req)
	})

	h.handle(getPath, func(req Request) Response {
		return h.driver.Get(req)
	})

	h.handle(listPath, func(req Request) Response {
		return h.driver.List(req)
	})

	h.handle(removePath, func(req Request) Response {
		return h.driver.Remove(req)
	})

	h.handle(hostVirtualPath, func(req Request) Response {
		return h.driver.Path(req)
	})

	h.handleMount(mountPath, func(req MountRequest) Response {
		return h.driver.Mount(req)
	})

	h.handleUnmount(unmountPath, func(req UnmountRequest) Response {
		return h.driver.Unmount(req)
	})
	h.handle(capabilitiesPath, func(req Request) Response {
		return h.driver.Capabilities(req)
	})
}

// HandleFunc registers a function to handle a request path with.
func (h Handler) HandleFunc(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	h.mux.HandleFunc(path, fn)
}

//Rewrite standart func (w htt.ResponseWriter, r *http.Request) to func (r Request): Response
func (h *Handler) handle(name string, actionCall actionHandler) {
	h.HandleFunc(name, func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := DecodeRequest(w, r, &req); err != nil {
			return
		}

		res := actionCall(req)

		EncodeResponse(w, res, res.Err)
	})
}

//Rewrite and extend func (w htt.ResponseWriter, r *http.Request) to func (r Request): Response
func (h *Handler) handleMount(name string, actionCall mountActionHandler) {
	h.HandleFunc(name, func(w http.ResponseWriter, r *http.Request) {
		var req MountRequest
		if err := DecodeRequest(w, r, &req); err != nil {
			return
		}

		res := actionCall(req)
		EncodeResponse(w, res, res.Err)
	})
}

//Rewrite and extend func (w htt.ResponseWriter, r *http.Request) to func (r Request): Response
func (h *Handler) handleUnmount(name string, actionCall unmountActionHandler) {
	h.HandleFunc(name, func(w http.ResponseWriter, r *http.Request) {
		var req UnmountRequest
		if err := DecodeRequest(w, r, &req); err != nil {
			return
		}

		res := actionCall(req)
		EncodeResponse(w, res, res.Err)
	})
}
