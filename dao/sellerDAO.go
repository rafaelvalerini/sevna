package dao

import (
	"cevafacil.com.br/model"
	"cevafacil.com.br/service"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sort"
)

func FindSellerByIdAndSize(id int, size int, lat float64, lng float64) (seller model.Sellers) {

	session, err := mgo.Dial("mongodb://admin:admin@ds015720.mlab.com:15720/cevafacil")
	//session, err := mgo.Dial("mongodb://localhost:27017/cevafacil")

	if err != nil {

		panic(err)

	}

	defer session.Close()

	c := session.DB("cevafacil").C("sellers")

	var beers model.Sellers

	err = c.Find(bson.M{"beers.beer": id, "beers.size": size}).All(&beers)

	if err != nil {

		panic(err)

	}

	if beers != nil {

		for i := 0; i < len(beers); i++ {

			beers[i].Distance = fmt.Sprintf("%.2fKm", service.Distance(lat, lng, beers[i].Lat, beers[i].Lng)/1000)

			beers[i].LatLocal = lat

			beers[i].LngLocal = lng

			var beersNil = make([]model.Beer, 0)

			var beersAux = make([]model.Beer, 1)

			for j := 0; j < len(beers[i].Beers); j++ {

				if beers[i].Beers[j].Size == size && beers[i].Beers[j].Beer == id {

					beersAux[0] = beers[i].Beers[j]

				}

			}

			beers[i].Beers = beersNil

			beers[i].Beer = beersAux[0]

		}

	}

	sort.Sort(beers)

	return beers

}

func FindSellerById(id int, size int, lat float64, lng float64) (seller model.Seller) {

	session, err := mgo.Dial("mongodb://admin:admin@ds015720.mlab.com:15720/cevafacil")
	//session, err := mgo.Dial("mongodb://localhost:27017/cevafacil")

	if err != nil {

		panic(err)

	}

	defer session.Close()

	c := session.DB("cevafacil").C("sellers")

	var beer model.Seller

	err = c.Find(bson.M{"id": id}).One(&beer)

	if err != nil {

		panic(err)

	}

	beer.Distance = fmt.Sprintf("%.2fKm", service.Distance(lat, lng, beer.Lat, beer.Lng)/1000)

	var beersNil = make([]model.Beer, 0)

	var beersAux = make([]model.Beer, 1)

	for j := 0; j < len(beer.Beers); j++ {

		if beer.Beers[j].Size == size && beer.Beers[j].Beer == id {

			beersAux[0] = beer.Beers[j]

		}

	}

	beer.Beers = beersNil

	beer.Beer = beersAux[0]

	return beer

}
