package handler

type Handler interface {
	Register(*Handler) error
	Fire(*Event) error
}

type handlers map[string][]*Handler

func (h *Handler) Register(n *handler) error {
	for _, level := range n.Levels() {
		handlers[level] = append(handlers[n.Name], n)
	}
}

func (h *Handler) Fire(alert *Alert) error {

}
