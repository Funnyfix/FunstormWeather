package owmhelper

import (
	"fmt"
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
		return "Неопознаная локация"
	}
	text := fmt.Sprintf("На улице %s \nТемпература: %.f℃ \nОщущается как: %.f℃ \nВетер: %.2f м/c", data.Weather[0].Description, data.Main.Temp, data.Main.FeelsLike, data.Wind.Speed)
	return text
}
