package main

import (
	"github.com/seniorescobar/design-patterns/cmd/observer/alertsystem"
	"github.com/seniorescobar/design-patterns/cmd/observer/logger"
	"github.com/seniorescobar/design-patterns/cmd/observer/ui"
	"github.com/seniorescobar/design-patterns/cmd/observer/weatherstation"
)

func main() {
	// init observers
	ui := &ui.UserInterface{}
	logger := &logger.Logger{}
	alertsystem := &alertsystem.AlertSystem{}

	// register observers
	ws := &weatherstation.WeatherStation{}
	ws.RegisterObserver(ui)
	ws.RegisterObserver(logger)
	ws.RegisterObserver(alertsystem)

	// set weather data
	ws.SetWeatherData(&weatherstation.WeatherData{
		Temperature: 32,
		WindSpeed:   20,
		Pressure:    33,
	})
	ws.SetWeatherData(&weatherstation.WeatherData{
		Temperature: 83,
		WindSpeed:   35,
		Pressure:    76,
	})
}
