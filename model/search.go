package model

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
}

type Sellers []Seller

type Beer struct {
	Beer  int
	Size  int
	Value string
}

type Beers []Beer
