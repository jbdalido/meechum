package meechum

type Handler interface {
	Fire(*Result) error
	Levels() []ErrorCode
	String() string
}

var handlers map[string][]Handler

func RegisterHandler(n Handler) error {
	if handlers == nil {
		handlers = make(map[string][]Handler, 50)
	}
	for _, level := range n.Levels() {
		handlers[level.String()] = append(handlers[n.String()], n)
	}
	return nil
}
