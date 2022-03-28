package weatherstation

type Observer interface {
	Update(*WeatherData)
}

type (
	WeatherData struct {
		Temperature, WindSpeed, Pressure float32
	}

	WeatherStation struct {
		observers []Observer
	}
)

func (ws *WeatherStation) RegisterObserver(observer Observer) {
	ws.observers = append(ws.observers, observer)
}

func (ws *WeatherStation) SetWeatherData(wd *WeatherData) {
	ws.notifyObservers(wd)
}

func (ws *WeatherStation) notifyObservers(wd *WeatherData) {
	for _, o := range ws.observers {
		o.Update(wd)
	}
}
