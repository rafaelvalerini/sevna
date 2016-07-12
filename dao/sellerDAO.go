package dao

import (
	"delivery.futuroclick.com.br/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"os/exec"
	"strings"
)

func AuthSeller(user string, password string) (meta model.Seller) {

	session, err := mgo.Dial("mongodb://delivery:delivery_money2@ds011745.mlab.com:11745/delivery_app_futuroclick")
	//session, err := mgo.Dial("mongodb://localhost:27017/cevafacil")

	if err != nil {

		panic(err)

	}

	defer session.Close()

	c := session.DB("delivery_app_futuroclick").C("sellers")

	var seller model.Seller

	err = c.Find(bson.M{"user" : user, "password" : password}).One(&seller)

	if err != nil {

		panic(err)

	}

	return seller

}

func AuthConsumer(id int64, user string, password string) (meta model.SellerUser) {

	session, err := mgo.Dial("mongodb://delivery:delivery_money2@ds011745.mlab.com:11745/delivery_app_futuroclick")
	//session, err := mgo.Dial("mongodb://localhost:27017/cevafacil")

	if err != nil {

		panic(err)

	}

	defer session.Close()

	c := session.DB("delivery_app_futuroclick").C("sellers_consumers")

	var seller model.SellerUser

	err = c.Find(bson.M{"id_seller": id, "user": user, "password" : password}).One(&seller)

	if err != nil {

		panic(err)

	}

	return seller

}

func SaveOrder(id int64, entity model.SellerOrder) (meta model.SellerOrder) {

	session, err := mgo.Dial("mongodb://delivery:delivery_money2@ds011745.mlab.com:11745/delivery_app_futuroclick")
	//session, err := mgo.Dial("mongodb://localhost:27017/cevafacil")

	if err != nil {

		panic(err)

	}

	defer session.Close()

	c := session.DB("delivery_app_futuroclick").C("sellers_orders")

	var sellerOrder model.SellerOrder

	err = c.Find(bson.M{"id_seller": id}).Sort("-id").One(&sellerOrder)

	if err != nil {

		entity.Id = int64(1);

	}else{

		entity.Id = sellerOrder.Id + 1

	}

	entity.IdSeller = id

	entity.DateCreate = time.Now()

	timeSession := session.DB("delivery_app_futuroclick").C("sellers_time_delivery")

	var timeDelivery model.TimeToDelivery

	err = timeSession.Find(bson.M{"id_seller": id}).One(&timeDelivery)

	if err != nil {

		panic(err)

	}

	entity.TimeToDelivery = timeDelivery

	err = c.Insert(entity)

	if err != nil {

		panic(err)

	}

	
	return entity

}

func GetStoresBySeller(seller int64) (stores []model.SellerStore) {

	session, err := mgo.Dial("mongodb://delivery:delivery_money2@ds011745.mlab.com:11745/delivery_app_futuroclick")
	//session, err := mgo.Dial("mongodb://localhost:27017/cevafacil")

	if err != nil {

		panic(err)

	}

	defer session.Close()

	c := session.DB("delivery_app_futuroclick").C("sellers_stores")

	var sellers []model.SellerStore

	err = c.Find(bson.M{"id_seller": seller}).All(&seller)

	if err != nil {

		panic(err)

	}

	return sellers

}

func GetOrdersByConsumer(seller int64, consumer string) (stores []model.SellerOrder) {

	session, err := mgo.Dial("mongodb://delivery:delivery_money2@ds011745.mlab.com:11745/delivery_app_futuroclick")
	//session, err := mgo.Dial("mongodb://localhost:27017/cevafacil")

	if err != nil {

		panic(err)

	}

	defer session.Close()

	c := session.DB("delivery_app_futuroclick").C("sellers_orders")

	var sellers []model.SellerOrder

	err = c.Find(bson.M{"id_seller": seller, "id_user": consumer}).All(&sellers)

	if err != nil {

		panic(err)

	}

	return sellers

}

func GetOrdersBySeller(seller int64, status int, store int, order int64, limit int) (stores []model.SellerOrder) {

	session, err := mgo.Dial("mongodb://delivery:delivery_money2@ds011745.mlab.com:11745/delivery_app_futuroclick")
	//session, err := mgo.Dial("mongodb://localhost:27017/cevafacil")

	if err != nil {

		panic(err)

	}

	defer session.Close()

	c := session.DB("delivery_app_futuroclick").C("sellers_orders")

	var sellers []model.SellerOrder

	conditions := make(bson.M, 0)

	conditions["id_seller"] = seller

	if status != 0 {
		conditions["status"] = status
	}
	
	if store != 0 {
		conditions["id_store"] = store
	}

	if order != 0 {
		conditions["id"] = order
	}
	
	err = c.Find(conditions).Limit(limit).Sort("-date_create").All(&sellers)

	if err != nil {

		panic(err)

	}

	return sellers

}


func SaveUser(id int64, entity model.SellerUser) (meta model.Meta) {

	session, err := mgo.Dial("mongodb://delivery:delivery_money2@ds011745.mlab.com:11745/delivery_app_futuroclick")
	//session, err := mgo.Dial("mongodb://localhost:27017/cevafacil")

	if err != nil {

		panic(err)

	}

	defer session.Close()

	c := session.DB("delivery_app_futuroclick").C("sellers_consumers")

	out, err := exec.Command("uuidgen").Output()

	entity.Id = string(out[:])

	entity.Id = strings.Replace(entity.Id,"\n","",-1)

	entity.IdSeller = id

	entity.DateCreate = time.Now()

	err = c.Insert(entity)

	if err != nil {

		panic(err)

	}

	entry := model.Meta{Value: entity.Id}

	return entry

}

func UpdateStatusOrder(id int64, user string, order int64, status int) (meta model.Meta) {

	session, err := mgo.Dial("mongodb://admin:admin@ds015720.mlab.com:15720/cevafacil")
	//session, err := mgo.Dial("mongodb://localhost:27017/cevafacil")

	if err != nil {

		panic(err)

	}

	defer session.Close()

	c := session.DB("cevafacil").C("sellers")

	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"status": status}},
		ReturnNew: true,
	}

	var sellerOrder model.SellerOrder

	info, err := c.Find(bson.M{"_id": order, "id_user": user, "id_seller" : id}).Apply(change, &sellerOrder)

	if info.Updated != 1 {

		entry := model.Meta{Value: "NOK"}

		return entry
	}

	if err != nil {

		panic(err)

	}

	entry := model.Meta{Value: "OK"}

	return entry

}

func SaveTimeDelivery(id int64, time int) (meta model.Meta) {
	session, err := mgo.Dial("mongodb://admin:admin@ds015720.mlab.com:15720/cevafacil")
	//session, err := mgo.Dial("mongodb://localhost:27017/cevafacil")

	if err != nil {

		panic(err)

	}

	defer session.Close()

	c := session.DB("cevafacil").C("sellers_time_delivery")

	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"time_to_delivery": time}},
		ReturnNew: true,
	}

	var timeDelivery model.TimeToDelivery

	info, err := c.Find(bson.M{"id": id}).Apply(change, &timeDelivery)

	if info.Updated != 1 {

		entry := model.Meta{Value: "NOK"}

		return entry
	}

	if err != nil {

		panic(err)

	}

	entry := model.Meta{Value: "OK"}

	return entry
}