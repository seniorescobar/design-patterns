package printer

import "fmt"

type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("Printing from Epson printer")
}
