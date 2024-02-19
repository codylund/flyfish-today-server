package sites

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Site struct {
	ObjectID   primitive.ObjectID `json:"_id"         bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"-"           bson:"user_id"`
	SiteID     string             `json:"site_id"     bson:"site_id"`
	IsFavorite bool               `json:"is_favorite" bson:"is_favorite"`
	Tags       Tags               `json:"tags"        bson:"tags"`
}

type SiteUpdate struct {
	IsFavorite *bool     `json:"is_favorite,omitempty" bson:"is_favorite,omitempty"`
	Tags       *[]string `json:"tags,omitempty" bson:"tags,omitempty"`
}
