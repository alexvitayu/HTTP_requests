package main

import (
	"HTTP_requests/weather/geo"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Начинаем новый проект!")
	city := flag.String("city", "", "Город пользователя")

	flag.Parse()

	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)
}
