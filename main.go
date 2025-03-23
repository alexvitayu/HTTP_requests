package main

import (
	"HTTP_requests/weather/geo"
	"HTTP_requests/weather/weather"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Начинаем новый проект!")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
