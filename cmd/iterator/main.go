package main

import (
	"fmt"

	"github.com/seniorescobar/design-patterns/cmd/iterator/menu"
	"github.com/seniorescobar/design-patterns/cmd/iterator/menu/dinermenu"
)

func main() {
	var m menu.Menu = dinermenu.Default
	var mi menu.Iterator = m.CreateIterator()

	for mi.HasNext() {
		f := mi.Next()

		fmt.Println(f)
	}
}
