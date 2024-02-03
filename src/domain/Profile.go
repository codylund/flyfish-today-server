package domain

type Profile struct {
    Username    string `json:"username"     bson:"username"`
    DisplayName string `json:"display_name" bson:"display_name"`
}
