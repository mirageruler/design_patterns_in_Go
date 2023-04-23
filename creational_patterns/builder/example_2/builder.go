package main

const (
	HouseCondo     string = "Condo"
	HouseApartment        = "Apartment"
	GarageBig             = "Big"
	GarageSmall           = "Small"
)

type iBuilder interface {
	SetCategory(string) iBuilder
	SetWindow(int) iBuilder
	SetFloor(int) iBuilder
}
