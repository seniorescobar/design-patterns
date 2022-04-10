package computer

import (
	"fmt"

	"github.com/seniorescobar/design-patterns/cmd/bridge/printer"
)

type Windows struct {
	p printer.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request from Windows")
	w.p.PrintFile()
}

func (w *Windows) SetPrinter(p printer.Printer) {
	w.p = p
}
