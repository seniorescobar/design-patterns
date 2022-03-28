package ui

import (
	"fmt"

	"github.com/seniorescobar/design-patterns/cmd/observer/weatherstation"
)

type UserInterface struct {
	wd *weatherstation.WeatherData
}

func (ui *UserInterface) Update(wd *weatherstation.WeatherData) {
	ui.wd = wd
	ui.display()
}

func (ui *UserInterface) display() {
	fmt.Println("displaying: ", ui.wd)
}
