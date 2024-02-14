package domain

type User struct {
	Username    string `json:"username"     bson:"username"`
	DisplayName string `json:"display_name" bson:"display_name"`
	Password    string `json:"password"     bson:"password"`
}
