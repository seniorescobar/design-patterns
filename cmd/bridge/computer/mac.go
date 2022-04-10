package computer

import (
	"fmt"

	"github.com/seniorescobar/design-patterns/cmd/bridge/printer"
)

type Mac struct {
	p printer.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request from Mac")
	m.p.PrintFile()
}

func (m *Mac) SetPrinter(p printer.Printer) {
	m.p = p
}
