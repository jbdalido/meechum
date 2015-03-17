package api

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/jbdalido/meechum"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// HERE IS THE API
type Api struct {
	Listen  string
	API     *Api
	Router  *httprouter.Router
	Meechum *meechum.Runtime
}

type ApiError struct {
	err error `json:"error"`
}

func NewApi(meechum *meechum.Runtime, listen string) (*Api, error) {

	// TODO: Check listen correct

	router := httprouter.New()
	// Setup routes
	router.GET("/status", a.serveAlive)

	router.GET("/v1/nodes", a.serveListNodes)
	router.GET("/v1/node/{id}", a.serveStatusNode)
	router.POST("/v1/node", a.serveAddNode)
	router.PUT("/v1/node", a.serveUpdateNode)
	router.DELETE("/v1/node", a.serveUpdateNode)

	router.GET("/v1/checks", a.serveListChecks)
	router.GET("/v1/check/{id}", a.serveStatusCheck)
	router.POST("/v1/check", a.serveAddCheck)
	router.PUT("/v1/check", a.serveUpdateCheck)
	router.DELETE("/v1/check", a.serveUpdateCheck)

	router.GET("/v1/checks", a.serveListGroups)
	router.GET("/v1/group/{id}", a.serveStatusGroup)
	router.POST("/v1/group", a.serveAddGroup)
	router.PUT("/v1/group", a.serveUpdateGroup)
	router.DELETE("/v1/group", a.serveUpdateGroup)

	router.POST("/v1/alert", a.serveAlert)

	return &API{
		Listen:  listen,
		Router:  router,
		Meechum: meechum,
	}, nil
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
func (a *Api) serveListNodes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := a.Meechum.GetListNodes()
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	a.jsonResponse(data, w)
}

func (a *Api) serveStatusNodes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := a.Meechum.GetStatusNodes()
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	a.jsonResponse(data, w)
}

func (a *Api) serveStatusNode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data, err := a.Meechum.GetListNode(ps.ByName("id"))
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	a.jsonResponse(data, w)
}

func (a *Api) serveDeleteNode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := a.Meechum.DeleteNode(ps.ByName("id"))
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveUpdateNode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	node := &meechum.Node
	err := a.decodeRequest(r, node)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	data, err := a.Meechum.CreateNode(node)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveCreateNode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO : handle post request in a better way ?
	node := &meechum.Node
	err := a.decodeRequest(r, node)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	data, err := a.Meechum.CreateNode(node)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)

}

// Checks

func (a *Api) serveDeleteCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := a.Meechum.DeleteCheck(ps.ByName("id"))
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveUpdateCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO : handle post request in a better way ?
	decoder := json.NewDecoder(r.Body)
	node := &meechum.Node
	err := decoder.Decode(node)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	data, err := a.Meechum.UpdateNode(node)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveCreateCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO : handle post request in a better way ?
	check := &meechum.Check
	err := a.decodeRequest(r, check)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	data, err := a.Meechum.CreateCheck(check)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

// Groups

func (a *Api) serveDeleteGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := a.Meechum.DeleteGroup(ps.ByName("id"))
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveUpdateGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	group := &meechum.Group
	err := a.decodeRequest(r, group)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	data, err := a.Meechum.UpdateGroup(group)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) serveCreateGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	group := &meechum.Group
	err := a.decodeRequest(r, group)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}

	data, err := a.Meechum.CreateGroup(group)
	if err != nil {
		a.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Api) decodeRequest(r *http.Request, strc *interface{}) error {
	// TODO : handle post request in a better way ?
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(node)
	if err != nil {
		return err
	}
}

func (a *Api) jsonErrorResponse(err error, w http.ResponseWriter) {

	e := &ApiError{
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
	jdata, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jdata)
	return nil
}
