package owmhelper

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"log"
	"os"
)

var icons = map[string]string{
	"01d": "☀️",
	"01n": "🌛",
	"02d": "🌤",
	"02n": "🌤",
	"03d": "☁️",
	"03n": "☁️",
	"04d": "☁️",
	"04n": "☁️",
	"09d": "🌧",
	"09n": "🌧",
	"10d": "🌦",
	"10n": "🌦",
	"11d": "⛈",
	"11n": "⛈",
	"13d": "❄️",
	"13n": "❄️",
	"50d": "🌫",
	"50n": "🌫",
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
	text := fmt.Sprintf("It's %s outside %s \nTemperature: %d℃ \nFeels like: %d℃ \nWind speed: %.2f m/s", data.Weather[0].Description, icon, maintemp, feelslike, data.Wind.Speed)
	return textcity + text
}
