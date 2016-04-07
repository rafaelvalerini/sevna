package model

import (
	"cevafacil.com.br/service"
)

type Seller struct {
	Id          int
	Name        string
	Img         string
	Distance    string
	Description string
	Business    []string
	Lat         float64
	Lng         float64
	Cards       []string
	Address     string
	Value       string
	Beers       []Beer
	Beer        Beer
	LatLocal    float64
	LngLocal    float64
}

type Sellers []Seller

type Beer struct {
	Beer  int
	Size  int
	Value string
}

type Beers []Beer

func (slice Sellers) Len() int {
	return len(slice)
}

func (slice Sellers) Less(i, j int) bool {

	return (service.Distance(slice[i].Lat, slice[i].Lng, slice[i].LatLocal, slice[i].LngLocal) / 1000) < (service.Distance(slice[j].Lat, slice[j].Lng, slice[j].LatLocal, slice[j].LngLocal) / 1000)

}

func (slice Sellers) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
