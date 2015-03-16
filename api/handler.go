package api

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/jbdalido/meechum"
	"net/http"
)

// HERE IS THE API

// servePing returns a simple response to let the client know the server is running.
func (a *Api) serveAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// Nodes

func (a *Api) serveListNodes(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (a *Api) serveStatusNodes(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (a *Api) serveDeleteNode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (a *Api) serveUpdateNode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (a *Api) serveCreateNode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// Checks

func (a *Api) serveDeleteCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (a *Api) serveUpdateCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (a *Api) serveCreateCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// Groups

func (a *Api) serveDeleteGroup(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (a *Api) serveUpdateGroup(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (a *Api) serveCreateGroup(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

//

func (a *Api) unmarshallRequest(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Unmarshall if failed, return failedfuckcode
	next(rw, r)
}
