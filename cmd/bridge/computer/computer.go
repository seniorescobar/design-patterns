package computer

import "github.com/seniorescobar/design-patterns/cmd/bridge/printer"

type Computer interface {
	Print()
	SetPrinter(printer.Printer)
}
