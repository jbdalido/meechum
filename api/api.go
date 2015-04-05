package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jbdalido/meechum"
	"log"
	"net/http"
)

// HERE IS THE Httpd
type Api struct {
	Listen  string
	Httpd   *Api
	Router  *mux.Router
	Meechum *meechum.Runtime
}

type HttpdError struct {
	err error `json:"error"`
}

func NewApi(meechum *meechum.Runtime, listen string) (*Api, error) {

	if meechum == nil {
		return nil, fmt.Errorf("Runtime cant be nil")
	}

	// TODO: Check listen correct
	a := &Api{
		Listen:  listen,
		Router:  mux.NewRouter(),
		Meechum: meechum,
	}

	// Setup routes
	a.Router.HandleFunc("/status", a.serveAlive).Methods("GET")

	a.Router.HandleFunc("/v1/nodes", a.serveListNodes).Methods("GET")
	a.Router.HandleFunc("/v1/node/{id}", a.serveStatusNode).Methods("GET")
	a.Router.HandleFunc("/v1/node", a.serveCreateNode).Methods("POST")
	a.Router.HandleFunc("/v1/node", a.serveUpdateNode).Methods("PUT")
	a.Router.HandleFunc("/v1/node", a.serveUpdateNode).Methods("DELETE")

	a.Router.HandleFunc("/v1/checks", a.serveListChecks).Methods("GET")
	a.Router.HandleFunc("/v1/check/{id}", a.serveStatusCheck).Methods("GET")
	a.Router.HandleFunc("/v1/check", a.serveCreateCheck).Methods("POST")
	a.Router.HandleFunc("/v1/check", a.serveUpdateCheck).Methods("PUT")
	a.Router.HandleFunc("/v1/check", a.serveUpdateCheck).Methods("DELETE")

	a.Router.HandleFunc("/v1/group/{id}", a.serveStatusGroup).Methods("GET")
	a.Router.HandleFunc("/v1/group", a.serveCreateGroup).Methods("POST")
	a.Router.HandleFunc("/v1/group", a.serveUpdateGroup).Methods("PUT")
	a.Router.HandleFunc("/v1/group", a.serveUpdateGroup).Methods("DELETE")

	a.Router.HandleFunc("/v1/alert", a.serveCreateAlert).Methods("POST")

	return a, nil
}

func (a *Api) Start() error {
	err := http.ListenAndServe(a.Listen, a.Router)
	if err != nil {
		return err
	}
	return nil
}

// servePing returns a simple response to let the client know the server is running.
func (a *Api) serveAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Nodes
func (a *Api) serveListNodes(w http.ResponseWriter, r *http.Request) {
	data, err := a.Meechum.GetListNodes()
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	a.jsonResponse(data, w)
}

func (a *Api) serveListGroups(w http.ResponseWriter, r *http.Request) {
	data, err := a.Meechum.GetListNodes()
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	a.jsonResponse(data, w)
}

func (a *Api) serveStatusNode(w http.ResponseWriter, r *http.Request) {
	data, err := a.Meechum.GetStatusNode(mux.Vars(r)["id"])
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	a.jsonResponse(data, w)
}

func (a *Api) serveDeleteNode(w http.ResponseWriter, r *http.Request) {
	err := a.Meechum.DeleteNode(mux.Vars(r)["id"])
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveUpdateNode(w http.ResponseWriter, r *http.Request) {
	node := &meechum.Node{}
	err := a.decodeRequest(r, node)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	err = a.Meechum.CreateNode(node)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveCreateNode(w http.ResponseWriter, r *http.Request) {
	// TODO : handle post request in a better way ?
	node := meechum.Node{}
	err := a.decodeRequest(r, node)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	err = a.Meechum.CreateNode(&node)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)

}

// Checks

func (a *Api) serveDeleteCheck(w http.ResponseWriter, r *http.Request) {
	err := a.Meechum.DeleteCheck(mux.Vars(r)["id"])
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveUpdateCheck(w http.ResponseWriter, r *http.Request) {
	// TODO : handle post request in a better way ?
	decoder := json.NewDecoder(r.Body)
	node := meechum.Node{}
	err := decoder.Decode(node)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	err = a.Meechum.UpdateNode(&node)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveCreateCheck(w http.ResponseWriter, r *http.Request) {
	// TODO : handle post request in a better way ?
	check := meechum.Check{}
	err := a.decodeRequest(r, check)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	err = a.Meechum.CreateCheck(&check)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveListChecks(w http.ResponseWriter, r *http.Request) {
	data, err := a.Meechum.ListChecks()
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	a.jsonResponse(data, w)
}

func (a *Api) serveStatusCheck(w http.ResponseWriter, r *http.Request) {
	data, err := a.Meechum.ListChecks()
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	a.jsonResponse(data, w)
}

// Groups
func (a *Api) serveDeleteGroup(w http.ResponseWriter, r *http.Request) {
	err := a.Meechum.DeleteGroup(mux.Vars(r)["id"])
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveUpdateGroup(w http.ResponseWriter, r *http.Request) {
	group := &meechum.Group{}
	err := a.decodeRequest(r, group)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	err = a.Meechum.UpdateGroup(group)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveCreateGroup(w http.ResponseWriter, r *http.Request) {
	group := &meechum.Group{}
	err := a.decodeRequest(r, group)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	err = a.Meechum.CreateGroup(group)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveStatusGroup(w http.ResponseWriter, r *http.Request) {
	data, err := a.Meechum.ListChecks()
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	a.jsonResponse(data, w)
}

func (a *Api) serveCreateAlert(w http.ResponseWriter, r *http.Request) {
	// TODO : handle post request in a better way ?
	alert := &meechum.Alert{}
	err := a.decodeRequest(r, alert)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	err = a.Meechum.CreateAlert(alert)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)

}

func (a *Api) decodeRequest(r *http.Request, strc interface{}) error {
	// TODO : handle post request in a better way ?
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(strc)
	if err != nil {
		return err
	}
	return nil
}

func (a *Api) jsonErrorResponse(err error, w http.ResponseWriter) {

	e := &HttpdError{
		err: err,
	}

	jdata, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		http.Error(w, "Server Error", 500)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jdata)

}

func (a *Api) jsonResponse(data interface{}, w http.ResponseWriter) {
	// Marshall the response
	jdata, err := json.Marshal(data)
	if err != nil {
		log.Printf("err %s", err)
		a.jsonErrorResponse(err, w)
	}
	log.Printf("data %s %s", data, jdata)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jdata)

}
