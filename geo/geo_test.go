package geo_test

import (
	"HTTP_requests/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	//Arrange
	city := "Brest"
	expected := geo.GeoData{
		City: "Brest",
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

func TestGetMyLocationNoCity(t *testing.T) {
	//Arrange
	city := "Bresttt"
	//Act
	_, err := geo.GetMyLocation(city)
	//Assert
	if err != geo.ErrNoCity {
		t.Errorf("ожидали %v, получили %v", geo.ErrNoCity, err)
	}
}
