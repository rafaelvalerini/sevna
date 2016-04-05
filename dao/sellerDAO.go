package dao

import (
	"cevafacil.com.br/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func FindSellerByIdAndSize(id int, size int) (seller model.Sellers) {

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

	var sellers []int

	for key := range beers {
		
		sellers[]

	}

	//[]interface{}{
	//    bson.D{{"key2", 2}},
	//    bson.D{{"key3", 2}},
	//}

	c := session.DB("cevafacil").C("sellers")

	var result model.Sellers

	err = c.Find(bson.M{"id": id}).All(&result)

	if err != nil {

		panic(err)

	}

	return result

}
