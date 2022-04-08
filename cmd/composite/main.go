package main

import "github.com/seniorescobar/design-patterns/cmd/composite/menu"

func main() {
	dinMenu := menu.NewMenu("DINER MENU")
	dinMenu.Add(menu.NewItem("Soup of the day"))
	dinMenu.Add(menu.NewItem("Hot Dog"))

	dessMenu := menu.NewMenu("DESSERT MENU")
	dessMenu.Add(menu.NewItem("Apple Pie"))
	dessMenu.Add(menu.NewItem("Cheesecake"))

	phMenu := menu.NewMenu("PANCAKE HOUSE MENU")
	phMenu.Add(menu.NewItem("Regular Pancake Breakfast"))
	phMenu.Add(menu.NewItem("Blueberry Pancakes"))
	phMenu.Add(menu.NewItem("Waffles"))
	phMenu.Add(dessMenu)

	allMenus := menu.NewMenu("ALL MENUS")
	allMenus.Add(dinMenu)
	allMenus.Add(phMenu)

	allMenus.Print()
}
