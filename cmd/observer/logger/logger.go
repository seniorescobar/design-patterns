package logger

import (
	"fmt"

	"github.com/seniorescobar/design-patterns/cmd/observer/weatherstation"
)

type Logger struct {
	wd *weatherstation.WeatherData
}

func (l *Logger) Update(wd *weatherstation.WeatherData) {
	l.wd = wd
	l.log()
}

func (l *Logger) log() {
	fmt.Println("logging: ", l.wd)
}
