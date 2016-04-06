package dao

import (
	"cevafacil.com.br/model"
	"cevafacil.com.br/service"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func FindSellerByIdAndSize(id int, size int, lat float64, lng float64) (seller model.Sellers) {

	//session, err := mgo.Dial("mongodb://admin:admin@ds015720.mlab.com:15720/cevafacil")
	session, err := mgo.Dial("mongodb://localhost:27017/cevafacil")

	if err != nil {

		panic(err)

	}

	defer session.Close()

	c := session.DB("cevafacil").C("sellers")

	var beers model.Sellers

	err = c.Find(bson.M{"beers.beer": id, "beers.size": size}).Limit(20).All(&beers)

	if err != nil {

		panic(err)

	}

	if beers != nil {

		for i := 0; i < len(beers); i++ {

			beers[i].Distance = strconv.FormatFloat(service.Distance(lat, lng, beers[i].Lat, beers[i].Lng)/1000, 'f', 6, 64)

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

	return beers

}
