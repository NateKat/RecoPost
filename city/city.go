package city

import (
	"RecoPost/office"
)

type City struct {
	cityName   string
	totParcels int
	offices    []office.Office
}

func New(cityName string, totParcels int, offices []office.Office) City {
	c := City{cityName, totParcels, offices}
	return c
}
