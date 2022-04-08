package menu

import "fmt"

type Item struct {
	name string
}

func NewItem(name string) MenuComponent {
	return &Item{
		name: name,
	}
}

func (i *Item) Add(MenuComponent) {
	panic("operation not supported")
}

func (i *Item) Print() {
	fmt.Println("- ", i.name)
}
