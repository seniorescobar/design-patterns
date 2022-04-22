package originator

import "github.com/seniorescobar/design-patterns/cmd/memento/memento"

type Originator struct {
	State string
}

func (o *Originator) CreateMemento() *memento.Memento {
	return &memento.Memento{
		State: o.State,
	}
}

func (o *Originator) SetMemento(m *memento.Memento) {
	o.State = m.State
}
