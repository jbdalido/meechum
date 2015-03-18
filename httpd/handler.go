package httpd

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/jbdalido/meechum"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// HERE IS THE Httpd
type Httpd struct {
	Listen  string
	Httpd   *Httpd
	Router  *httprouter.Router
	Meechum *meechum.Runtime
}

type HttpdError struct {
	err error `json:"error"`
}

func NewHttpd(meechum *meechum.Runtime, listen string) (*Httpd, error) {

	// TODO: Check listen correct

	router := httprouter.New()
	// Setup routes
	router.GET("/status", h.serveAlive)

	router.GET("/v1/nodes", h.serveListNodes)
	router.GET("/v1/node/{id}", h.serveStatusNode)
	router.POST("/v1/node", h.serveAddNode)
	router.PUT("/v1/node", h.serveUpdateNode)
	router.DELETE("/v1/node", h.serveUpdateNode)

	router.GET("/v1/checks", h.serveListChecks)
	router.GET("/v1/check/{id}", h.serveStatusCheck)
	router.POST("/v1/check", h.serveAddCheck)
	router.PUT("/v1/check", h.serveUpdateCheck)
	router.DELETE("/v1/check", h.serveUpdateCheck)

	router.GET("/v1/checks", h.serveListGroups)
	router.GET("/v1/group/{id}", h.serveStatusGroup)
	router.POST("/v1/group", h.serveAddGroup)
	router.PUT("/v1/group", h.serveUpdateGroup)
	router.DELETE("/v1/group", h.serveUpdateGroup)

	router.POST("/v1/alert", h.serveAlert)

	return &Httpd{
		Listen:  listen,
		Router:  router,
		Meechum: meechum,
	}, nil
}

func (h *Httpd) Start() error {
	err := http.ListenAndServe(h.Listen, h.Router)
	if err != nil {
		return err
	}
	return nil
}

// servePing returns a simple response to let the client know the server is running.
func (h *Httpd) serveAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Nodes
func (h *Httpd) serveListNodes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := h.Meechum.GetListNodes()
	if err != nil {
		h.jsonErrorResponse(err, w)
	}
	h.jsonResponse(data, w)
}

func (h *Httpd) serveStatusNodes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := h.Meechum.GetStatusNodes()
	if err != nil {
		h.jsonErrorResponse(err, w)
	}
	h.jsonResponse(data, w)
}

func (h *Httpd) serveStatusNode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data, err := h.Meechum.GetListNode(ps.ByName("id"))
	if err != nil {
		h.jsonErrorResponse(err, w)
	}
	h.jsonResponse(data, w)
}

func (h *Httpd) serveDeleteNode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := h.Meechum.DeleteNode(ps.ByName("id"))
	if err != nil {
		h.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Httpd) serveUpdateNode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	node := &meechum.Node
	err := h.decodeRequest(r, node)
	if err != nil {
		h.jsonErrorResponse(err, w)
	}

	data, err := h.Meechum.CreateNode(node)
	if err != nil {
		h.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Httpd) serveCreateNode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO : handle post request in a better way ?
	node := &meechum.Node
	err := h.decodeRequest(r, node)
	if err != nil {
		h.jsonErrorResponse(err, w)
	}

	data, err := h.Meechum.CreateNode(node)
	if err != nil {
		h.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)

}

// Checks

func (h *Httpd) serveDeleteCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := h.Meechum.DeleteCheck(ps.ByName("id"))
	if err != nil {
		h.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Httpd) serveUpdateCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO : handle post request in a better way ?
	decoder := json.NewDecoder(r.Body)
	node := &meechum.Node
	err := decoder.Decode(node)
	if err != nil {
		h.jsonErrorResponse(err, w)
	}

	data, err := h.Meechum.UpdateNode(node)
	if err != nil {
		h.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Httpd) serveCreateCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO : handle post request in a better way ?
	check := &meechum.Check
	err := h.decodeRequest(r, check)
	if err != nil {
		h.jsonErrorResponse(err, w)
	}

	data, err := h.Meechum.CreateCheck(check)
	if err != nil {
		h.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

// Groups

func (h *Httpd) serveDeleteGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := h.Meechum.DeleteGroup(ps.ByName("id"))
	if err != nil {
		h.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Httpd) serveUpdateGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	group := &meechum.Group
	err := h.decodeRequest(r, group)
	if err != nil {
		h.jsonErrorResponse(err, w)
	}

	data, err := h.Meechum.UpdateGroup(group)
	if err != nil {
		h.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Httpd) serveCreateGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	group := &meechum.Group
	err := h.decodeRequest(r, group)
	if err != nil {
		h.jsonErrorResponse(err, w)
	}

	data, err := h.Meechum.CreateGroup(group)
	if err != nil {
		h.jsonErrorResponse(err, w)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Httpd) decodeRequest(r *http.Request, strc *interface{}) error {
	// TODO : handle post request in a better way ?
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(node)
	if err != nil {
		return err
	}
}

func (h *Httpd) jsonErrorResponse(err error, w http.ResponseWriter) {

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

func (h *Httpd) jsonResponse(data interface{}, w http.ResponseWriter) {
	// Marshall the response
	jdata, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jdata)
	return nil
}
