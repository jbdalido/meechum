package meechum

type Engine struct {
	Status *Stats
}

type Stats struct {
	Requests int64
	Failures int64

	LastSeenBackend time.Time

	Checks            int
	ChecksFailures    int
	KeepAliveFailures int
	PoolSize          int
}
