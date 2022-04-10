package printer

import "fmt"

type HP struct{}

func (hp *HP) PrintFile() {
	fmt.Println("Printing from HP printer")
}
