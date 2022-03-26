package dinermenu

import "github.com/seniorescobar/design-patterns/cmd/iterator/menu"

type (
	DinerMenu struct {
		foods []menu.Food
	}
)

var Default *DinerMenu = &DinerMenu{
	foods: []menu.Food{
		{Name: "Pizza", Price: 10},
		{Name: "Carbonara", Price: 7},
	},
}

func (m *DinerMenu) CreateIterator() menu.Iterator {
	return &MenuIterator{
		foods: m.foods,
	}
}
