package menu

import "fmt"

type Menu struct {
	name string

	components []MenuComponent
}

func NewMenu(name string) MenuComponent {
	return &Menu{
		name:       name,
		components: make([]MenuComponent, 0),
	}
}

func (m *Menu) Add(component MenuComponent) {
	m.components = append(m.components, component)
}

func (m *Menu) Print() {
	fmt.Println()
	fmt.Println(m.name)
	fmt.Println("----------")

	for _, c := range m.components {
		c.Print()
	}
}
