package gumball

type GumballMachine struct {
	sold                  State
	noQuarter, hasQuarter State

	currState State
}

func NewGumballMachine() *GumballMachine {
	gm := &GumballMachine{}

	gm.noQuarter = NewNoQuarterState(gm)
	gm.hasQuarter = NewHasQuarterState(gm)
	gm.sold = NewSoldState(gm)

	gm.currState = gm.noQuarter

	return gm
}

func (g *GumballMachine) InsertQuarter() {
	g.currState.InsertQuarter()
}

func (g *GumballMachine) EjectQuarter() {
	g.currState.EjectQuarter()
}

func (g *GumballMachine) TurnCrank() {
	g.currState.TurnCrank()
	g.currState.Dispense()
}

func (g *GumballMachine) setState(state State) {
	g.currState = state
}
