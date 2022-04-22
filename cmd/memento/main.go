package main

import (
	"fmt"

	"github.com/seniorescobar/design-patterns/cmd/memento/caretaker"
	"github.com/seniorescobar/design-patterns/cmd/memento/originator"
)

func main() {
	c := &caretaker.Caretaker{}

	o := &originator.Originator{State: "a"}
	fmt.Println(o)
	c.AddMemento(o.CreateMemento()) // State = "a"

	o.State = "b"
	fmt.Println(o)
	c.AddMemento(o.CreateMemento()) // State = "b"

	o.State = "c"
	fmt.Println(o)
	c.AddMemento(o.CreateMemento()) // State = "c"

	// Set state back to "b"

	o.SetMemento(c.GetMemento(1)) // State = "b"
	fmt.Println(o)
}
