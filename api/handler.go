package api

import (
	"errors"
	"net/http"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

type Mux interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	mux    Mux
	prefix string
}

func newHandler(mux goahttp.Muxer, prefix string) *handler {
	return &handler{
		mux:    mux,
		prefix: prefix,
	}
}

func (h *handler) correctRequestPath(request *http.Request) (err error) {
	path := request.URL.Path

	if strings.HasPrefix(path, h.prefix) {
		path = path[len(h.prefix):]
		request.URL.Path = path
		request.RequestURI = path
	} else {
		err = errors.New("incorrect path")
	}

	return
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.prefix != "" {
		if err := h.correctRequestPath(r); err != nil {
			http.NotFound(w, r)
			return
		}
	}
	h.mux.ServeHTTP(w, r)
}
