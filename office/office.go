package office

import (
	"RecoPost/parcel"
)

type Office struct {
	officeName  string
	maxParcelWt int
	minParcelWt int
	parcels     []parcel.Parcel
}

func New(officeName string, maxParcelWt int, minParcelWt int, parcellist []parcel.Parcel) Office {
	o := Office{officeName, maxParcelWt, minParcelWt, parcellist}
	return o
}
