package menu

type MenuComponent interface {
	Add(MenuComponent)
	Print()

	// skipping for this tutorial
	// Remove(MenuComponent)
}
