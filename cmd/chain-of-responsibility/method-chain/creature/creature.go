package creature

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func NewCreature(name string, att, def int) *Creature {
	return &Creature{
		Name:    name,
		Attack:  att,
		Defense: def,
	}
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}
