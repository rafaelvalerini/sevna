package model

type Seller struct {
	Id       int64  `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	User     string `json:"user" bson:"user"`
	Password string `json:"password" bson:"password"`
	Address  string `json:"address" bson:"address"`
	Phone    string `json:"phone" bson:"phone"`
}
