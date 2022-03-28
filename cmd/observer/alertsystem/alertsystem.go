package alertsystem

import (
	"fmt"

	"github.com/seniorescobar/design-patterns/cmd/observer/weatherstation"
)

type AlertSystem struct {
	wd *weatherstation.WeatherData
}

func (as *AlertSystem) Update(wd *weatherstation.WeatherData) {
	as.wd = wd
	as.alert()
}

func (as *AlertSystem) alert() {
	fmt.Println("alerting: ", as.wd)
}
