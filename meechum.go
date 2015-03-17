package meechum

import (
	"github.com/rcrowley/go-metrics"
	"time"
)

type Runtime struct {
	Status   *Stats
	Backend  *Backend
	Handlers []Handlers
	Stats    *Stats
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
