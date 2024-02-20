package model

type UserDetails struct {
	Name string `json:"name" bson:"name"`
	PhoneNo string `json:"phoneno" bson:"phoneno"`
	Address string `json:"address" bson:"address"`
	EmailId string `json:"emailid" bson:"emailid"`
}

