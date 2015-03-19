package meechum

import (
	"encoding/json"
	"fmt"
	"github.com/jbdalido/meechum/handler"
	//	"github.com/rcrowley/go-metrics"
	"log"
	"time"
)

type Runtime struct {
	Status   *Stats
	Backend  Backend
	Handlers []handler.Handler
	Stats    *Stats
	Checks   []*Check
	Err      error
}

type Stats struct {
	Requests          int64
	Failures          int64
	Checks            int
	ChecksFailures    int
	KeepAliveFailures int
	PoolSize          int
	LastSeenBackend   time.Time
}

type Alert struct {
	level int
}

// NewRuntime returns a backend connected runtime
func NewRuntime(backend string, host string) (*Runtime, error) {

	b, err := NewBackend(backend)
	if err != nil {
		return nil, err
	}
	// Connect to the backend
	err = b.Connect(host)
	if err != nil {
		return nil, err
	}

	return &Runtime{
		Backend: b,
	}, nil

}

// Subscribe is retrieving configuration from consul every n minutes
func (r *Runtime) Subscribe(groups []string) error {
	// Retrieve data from consul and create executor
	if len(groups) == 0 {
		return fmt.Errorf("Error subscribing to 0 group (funny guy)")
	}

	var checklist []string

	for _, group := range groups {
		c, err := r.getChecksFromGroup(group)
		if err != nil {
			log.Printf("Cant retrieve configurations for group %s", group)
		}
		checklist = append(checklist, c...)
	}

	err := r.updateChecksList(checklist)
	if err != nil {
		return err
	}

	return nil
}

func (r *Runtime) getChecksFromGroup(group string) ([]string, error) {
	data, err := r.Backend.GetKey("/groups/" + group)
	if err != nil {
		return nil, err
	}
	g := &Group{}
	err = json.Unmarshal(data, g)
	if err != nil {
		return nil, err
	}

	return g.Checks, nil
}

func (r *Runtime) updateChecksList(checkList []string) error {
	if len(checkList) == 0 {
		return fmt.Errorf("Checklist is empty, sad story...")
	}
	return nil
}

func (r *Runtime) Run() error {
	return nil
}

func (r *Runtime) GetListNodes() ([]*Node, error) {
	return nil, nil
}
func (r *Runtime) GetStatusNode(id string) (*Node, error) {
	return nil, nil
}

func (r *Runtime) DeleteNode(id string) error {
	return nil
}
func (r *Runtime) CreateNode(n *Node) error {
	return nil
}

func (r *Runtime) DeleteCheck(id string) error {
	return nil
}
func (r *Runtime) UpdateGroup(g *Group) error {
	return nil
}

func (r *Runtime) DeleteGroup(id string) error {
	return nil
}
func (r *Runtime) UpdateNode(n *Node) error {
	return nil
}
func (r *Runtime) CreateCheck(c *Check) error {
	return nil
}
func (r *Runtime) CreateGroup(g *Group) error {
	return nil
}
func (r *Runtime) CreateAlert(a *Alert) error {
	return nil
}
func (r *Runtime) ListChecks() ([]*Check, error) {
	return nil, nil
}
