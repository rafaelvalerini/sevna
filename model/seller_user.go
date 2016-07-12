package model

import (
	"time"
)

type SellerUser struct {
	Id         string    `json:"id" bson:"id"`
	Name       string    `json:"name" bson:"name"`
	User       string    `json:"user" bson:"user"`
	Password   string    `json:"password" bson:"password"`
	Address    string    `json:"address" bson:"address"`
	Phone      string    `json:"phone" bson:"phone"`
	IdSeller   int64     `json:"id_seller" bson:"id_seller"`
	DateCreate time.Time `json:"date_create" bson:"date_create"`
}

type SellerOrder struct {
	Id             int64          `json:"id" bson:"id"`
	Address        string         `json:"address" bson:"address"`
	Phone          string         `json:"phone" bson:"phone"`
	TypeDelivery   int64          `json:"type_delivery" bson:"type_delivery"`
	IdSeller       int64          `json:"id_seller" bson:"id_seller"`
	User     SellerUser          `json:"user" bson:"user"`
	IdStore        int64          `json:"id_store" bson:"id_store"`
	Products       []Product      `json:"products" bson:"products"`
	TotalValue     string         `json:"total_value" bson:"total_value"`
	Status         int64          `json:"status" bson:"status"`
	DateCreate     time.Time      `json:"date_create" bson:"date_create"`
	TimeToDelivery TimeToDelivery `json:"time_to_delivery" bson:"time_to_delivery"`
}

type Product struct {
	Id         int64  `json:"id" bson:"id"`
	Name 		string `json:"name" bson:"name"`
	Menu       Menu    `json:"menu" bson:"menu"`
	Quantity   int    `json:"quantity" bson:"quantity"`
	ValueTotal string `json:"value_total" bson:"value_total"`
}

type Menu struct{
	Id int `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type TimeToDelivery struct {
	IdSeller       int64 `json:"id_seller" bson:"id_seller"`
	TimeToDelivery int   `json:"time_to_delivery" bson:"time_to_delivery"`
}
