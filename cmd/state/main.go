package main

import "github.com/seniorescobar/design-patterns/cmd/state/gumball"

func main() {
	gm := gumball.NewGumballMachine()

	// success
	gm.InsertQuarter()
	gm.EjectQuarter()
	gm.InsertQuarter()
	gm.TurnCrank()

	// no quarter
	gm.TurnCrank()
}
