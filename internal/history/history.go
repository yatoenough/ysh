package history

type History struct {
	executed []string
}

func NewHistory() *History {
	return &History{
		executed: []string{},
	}
}

func (h *History) Append(cmd string) {
	h.executed = append(h.executed, cmd)
}

func (h *History) Get() []string {
	return h.executed
}
