package geo_test

import (
	"HTTP_requests/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	//Arrange
	city := "Minsk"
	expected := geo.GeoData{
		City: "Minsk",
	}
	//Act
	got, err := geo.GetMyLocation(city)
	//Assert
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("ожидали %v, получили %v", got.City, expected.City)
	}
}
