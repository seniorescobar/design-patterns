package gumball

import "log"

type NoQuarter struct {
	gm *GumballMachine
}

func NewNoQuarterState(gm *GumballMachine) State {
	return &NoQuarter{
		gm: gm,
	}
}

func (s *NoQuarter) InsertQuarter() {
	log.Println("inserting quarter")
	s.gm.setState(s.gm.hasQuarter)
}

func (s *NoQuarter) EjectQuarter() {
	log.Println("no quarter to eject")
}

func (s *NoQuarter) TurnCrank() {
	log.Println("cannot turn crank without a quarter")
}

func (s *NoQuarter) Dispense() {
	log.Println("cannot dispense until crank is turned")
}
