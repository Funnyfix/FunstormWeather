package owmhelper

import (
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
)

func Connect() *owm.CurrentWeatherData {
	w, err := owm.NewCurrent("C", "ru", os.Getenv("OWM_API_KEY"))
	if err != nil {
		log.Fatalln(err)
	}
	return w
}

func CheckWeather(lat, long float64) *owm.CurrentWeatherData {
	w := Connect()
	w.CurrentByCoordinates(
		&owm.Coordinates{
			Longitude: long,
			Latitude:  lat,
		},
	)
	return w
}
