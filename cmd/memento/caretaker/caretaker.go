package caretaker

import "github.com/seniorescobar/design-patterns/cmd/memento/memento"

type Caretaker struct {
	memento []*memento.Memento
}

func (c *Caretaker) AddMemento(m *memento.Memento) {
	c.memento = append(c.memento, m)
}

func (c *Caretaker) GetMemento(i int) *memento.Memento {
	return c.memento[i]
}
