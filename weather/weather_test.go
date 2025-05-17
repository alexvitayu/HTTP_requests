package weather_test

import (
	"HTTP_requests/weather/geo"
	"HTTP_requests/weather/weather"
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
	got := weather.GetWeather(geo, format)
	//Assert
	if !strings.Contains(got, expected) {
		t.Errorf("ожидали %v, получили %v", expected, got)
	}

}
