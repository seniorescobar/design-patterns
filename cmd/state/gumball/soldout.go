package gumball

import "log"

type SoldOut struct{}

func (s *SoldOut) InsertQuarter() { log.Println("sold out") }
func (s *SoldOut) EjectQuarter()  { log.Println("sold out") }
func (s *SoldOut) TurnCrank()     { log.Println("sold out") }
func (s *SoldOut) Dispense()      { log.Println("sold out") }
