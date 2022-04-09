package gumball

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
}
