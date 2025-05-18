package weather_test

import (
	"HTTP_requests/weather/geo"
	"HTTP_requests/weather/weather"
	"fmt"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	// Arrange
	expected := "Brest"
	geo := geo.GeoData{
		City: expected,
	}
	format := 3
	//Act
	got, _ := weather.GetWeather(geo, format)
	//Assert
	if !strings.Contains(got, expected) {
		t.Errorf("ожидали %v, получили %v", expected, got)
	}

}

func TestGetWeatherWrongFormat(t *testing.T) {
	//Arrange
	expected := "Brest"
	geo := geo.GeoData{
		City: expected,
	}
	format := 5
	//Act
	got, _ := weather.GetWeather(geo, format)
	//Assert
	if !strings.Contains(got, fmt.Sprint(format)) {
		t.Errorf("ожидали %v, получили %v", expected, got)
	}
}

func TestGetWeatherWrongFormatTheSecond(t *testing.T) {
	//Arrange
	city := "Brest"
	geo := geo.GeoData{
		City: city,
	}
	format := -1
	//Act
	_, err := weather.GetWeather(geo, format)
	//Assert
	if err != weather.ErrWrongFormat {
		t.Errorf("ожидали %v, получили %v", weather.ErrWrongFormat, err)

	}

}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 115},
	{name: "0 format", format: 0},
	{name: "Minus format", format: -1},
}

func TestGetWeatherWrongFormatGroup(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "Brest"
			geo := geo.GeoData{
				City: expected,
			}
			//Act
			_, err := weather.GetWeather(geo, tc.format)
			//Assert
			if err != weather.ErrWrongFormat {
				t.Errorf("ожидали %v, получили %v", weather.ErrWrongFormat, err)

			}
		})
	}
}
