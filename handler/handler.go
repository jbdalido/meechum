package handler

type Handler interface {
	Register(interface{}) error
	Fire(interface{}) error
	Levels() []string
	String() string
}

var handlers map[string][]Handler

func Register(n Handler) error {
	for _, level := range n.Levels() {
		handlers[level] = append(handlers[n.String()], n)
	}
	return nil
}
