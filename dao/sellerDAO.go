package dao

import (
	"cevafacil.com.br/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"cevafacil.com.br/service"
	"strconv"
)

func FindSellerByIdAndSize(id int, size int, lat float64, lng float64) (seller model.Sellers) {

	session, err := mgo.Dial("mongodb://admin:admin@ds015720.mlab.com:15720/cevafacil")

	if err != nil {

		panic(err)

	}

	defer session.Close()

	c := session.DB("cevafacil").C("products_values")

	var beers model.Beers

	err = c.Find(bson.M{"beer": id, "size": size}).All(&beers)

	if err != nil {

		panic(err)

	}

	var sellers = make([]int,len(beers))

	for key := range beers {
		
		sellers[key] = beers[key].Seller

	}

	var result model.Sellers

	d := session.DB("cevafacil").C("sellers")

	query := bson.M{"id": bson.M{"$in": sellers}}

	err = d.Find(query).All(&result)

	if err != nil {

		panic(err)

	}

	if result != nil {
		
		for i := 0; i < len(result); i++ {
			
			for j := 0; j < len(beers); j++ {
				
				if result[i].Id == beers[j].Seller {

					result[i].Value = beers[j].Value

					result[i].Distance = strconv.FormatFloat(service.Distance(lat, lng, result[i].Lat, result[i].Lng), 'f', 6, 64)

				}

			}

		}

	}

	return result

}
