package main

import (
	"github.com/seniorescobar/design-patterns/cmd/bridge/computer"
	"github.com/seniorescobar/design-patterns/cmd/bridge/printer"
)

func main() {
	hp := &printer.HP{}
	epson := &printer.Epson{}

	win := &computer.Windows{}
	win.SetPrinter(hp)
	win.Print()

	win.SetPrinter(epson)
	win.Print()

	mac := &computer.Mac{}
	mac.SetPrinter(hp)
	mac.Print()

	mac.SetPrinter(epson)
	mac.Print()
}
