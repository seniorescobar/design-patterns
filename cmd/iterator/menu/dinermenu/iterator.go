package dinermenu

import "github.com/seniorescobar/design-patterns/cmd/iterator/menu"

type MenuIterator struct {
	position int
	foods    []menu.Food
}

func (i *MenuIterator) HasNext() bool {
	if i.position < len(i.foods) {
		return true
	}

	return false
}

func (i *MenuIterator) Next() menu.Food {
	if i.HasNext() {
		f := i.foods[i.position]
		i.position++
		return f
	}

	return menu.Food{}
}
