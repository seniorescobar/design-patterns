package gumball

import "log"

type Sold struct {
	gm *GumballMachine
}

func NewSoldState(gm *GumballMachine) State {
	return &Sold{
		gm: gm,
	}
}

func (s *Sold) InsertQuarter() {
	log.Println("gumball already sold")
}

func (s *Sold) EjectQuarter() {
	log.Println("gumball already sold")
}

func (s *Sold) TurnCrank() {
	log.Println("gumball already sold")
}

func (s *Sold) Dispense() {
	log.Println("dispensing gumball")
	s.gm.setState(s.gm.noQuarter)
}
