package meechum

import (
	"encoding/json"
	"fmt"
	"github.com/jbdalido/meechum/handler"
	//	"github.com/rcrowley/go-metrics"
	"time"
)

type Runtime struct {
	Status   *Stats
	Backend  *Backend
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
	}

}

// Subscribe is retrieving configuration from consul every n minutes
func (r *Runtime) Subscribe(groups []string) error {
	// Retrieve data from consul and create executor
	if len(groups) == 0 {
		return fmt.Errorf("Error subscribing to 0 group (funny guy)")
	}

	for _, group := range groups {
		c = r.getChecksFromGroup(group)
		if err != nil {
			log.Printf("Cant retrieve configurations for group %s", group)
		}
		c.Checks = append(c.Checks, c...)
	}

	err := r.UpdateCheckList()
	if err != nil {
		return err
	}

	return nil
}

func (r *Runtime) getChecksFromGroup(group string) ([]Check, error) {
	data, err := r.Backend.GetKey("/groups/" + group)
	if err != nil {
		return err
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
}
