package owmhelper

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"log"
	"os"
)

var icons = map[string]string{
	"01d": "āļø",
	"01n": "š",
	"02d": "š¤",
	"02n": "š¤",
	"03d": "āļø",
	"03n": "āļø",
	"04d": "āļø",
	"04n": "āļø",
	"09d": "š§",
	"09n": "š§",
	"10d": "š¦",
	"10n": "š¦",
	"11d": "ā",
	"11n": "ā",
	"13d": "āļø",
	"13n": "āļø",
	"50d": "š«",
	"50n": "š«",
}

func Connect() *owm.CurrentWeatherData {
	w, err := owm.NewCurrent("C", "en", os.Getenv("OWM_API_KEY"))
	if err != nil {
		log.Fatalln(err)
	}
	return w
}

func CurrentWeatherByCoordinates(lat, long float64) *owm.CurrentWeatherData {
	w := Connect()
	w.CurrentByCoordinates(
		&owm.Coordinates{
			Longitude: long,
			Latitude:  lat,
		},
	)
	log.Println(w)
	return w
}

func CurrentWeatherByName(place string) *owm.CurrentWeatherData {
	w := Connect()
	w.CurrentByName(place)
	log.Println(w)
	return w
}

func ParseWeather(data *owm.CurrentWeatherData) string {
	if len(data.Weather) == 0 {
		return "Unknown location"
	}
	maintemp := int(data.Main.Temp)
	feelslike := int(data.Main.FeelsLike)
	icon := icons[data.Weather[0].Icon]
	textcity := fmt.Sprintf("%s, %s\n", data.Name, data.Sys.Country)
	text := fmt.Sprintf("It's %s outside %s \nTemperature: %dā \nFeels like: %dā \nWind speed: %.2f m/s", data.Weather[0].Description, icon, maintemp, feelslike, data.Wind.Speed)
	return textcity + text
}
