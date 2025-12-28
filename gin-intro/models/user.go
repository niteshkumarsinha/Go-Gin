package models


type Address struct {
	State string `json:"state" bson:"state"`
	City string `json:"city" bson:"city"`
	Country string `json:"country" bson:"country"`
	Pincode int `json:"pincode" bson:"pincode"`
}

type User struct {
	Name string `json:"name" bson:"user_name"`
	Email string `json:"email" bson:"user_email"`
	Address Address `json:"address" bson:"user_address"`
}

