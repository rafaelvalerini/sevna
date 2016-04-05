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
}

type Sellers []Seller

type Beer struct {
	Seller int
	Beer   int
	Size   int
	Value  string
}

type Beers []Beer
