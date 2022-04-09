package gumball

import "log"

type HasQuarter struct {
	gm *GumballMachine
}

func NewHasQuarterState(gm *GumballMachine) State {
	return &HasQuarter{
		gm: gm,
	}
}

func (s *HasQuarter) InsertQuarter() {
	log.Println("quarter already inserted")
}

func (s *HasQuarter) EjectQuarter() {
	log.Println("ejecting quarter")
	s.gm.setState(s.gm.noQuarter)
}

func (s *HasQuarter) TurnCrank() {
	log.Println("turning crank")
	s.gm.setState(s.gm.sold)
}

func (s *HasQuarter) Dispense() {
	log.Println("cannot dispense until crank is turned")
}
