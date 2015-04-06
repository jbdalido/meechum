package meechum

import (
	"encoding/json"
	"fmt"
	//	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"os/signal"
	"time"
)

//go:generate stringer -type=ErrorCode
type ErrorCode uint

const (
	OK ErrorCode = iota
	WARNING
	FATAL
)

type Runtime struct {
	Status  *Stats            `json:""`
	Backend Backend           `json:""`
	Stats   *Stats            `json:"stats"`
	Checks  map[string]*Check `json:"checks"`
	Result  chan *Result      `json:""`
	Err     error             `json:""`
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

	b, err := NewBackend(backend, host)
	if err != nil {
		return nil, err
	}

	return &Runtime{
		Backend: b,
		Checks:  make(map[string]*Check, 150),
		Result:  make(chan *Result, 5000),
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
		log.Printf("[Engine] Retrieving checks for group %s", group)
		c, err := r.getChecksFromGroup(group)
		if err != nil {
			log.Printf("Cant retrieve configurations for group %s", group)
		} else {
			log.Printf("[Engine] Checks for group %s retrieved", c)
			checklist = append(checklist, c...)
		}
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
	if data == nil {
		return nil, fmt.Errorf("Group %s is empty", group)
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

	for _, check := range checkList {

		data, err := r.Backend.GetKey("/checks/" + check)
		if err != nil {
			log.Printf("[Engine] Check %s does not exist %s", check, err)
			continue
		}

		c, err := NewCheck(data, r.Result)
		if err != nil {
			log.Printf("[Engine] Check %s creation error %s", check, err)
		}

		r.Checks[check] = c
	}
	return nil
}

func (r *Runtime) Run() error {
	killChannel := make(chan os.Signal, 1)
	signal.Notify(killChannel, os.Interrupt)

	for name, c := range r.Checks {
		log.Printf("Running check %s", name)
		go c.Run()
	}

	for {
		select {
		case d := <-r.Result:
			log.Printf("[Engine] Handling error [%s] LEVEL:%s", r.Result, d.Level)
			for _, handler := range handlers[d.Level] {
				go func() {
					err := handler.Fire(d)
					if err != nil {
						log.Printf("[Engine] Handler \"%s\" failed to fire event %s", handler, r.Result)
					}
				}()
			}
		case <-killChannel:
			log.Fatalf("Stop")
		}
	}
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
func (r *Runtime) ListChecks() (map[string]*Check, error) {
	return r.Checks, nil
}

func (r *Runtime) LiveStats() (string, error) {
	return fmt.Sprintf("Error Pool : %d", len(r.Result)), nil

}
