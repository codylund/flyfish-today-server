package domain

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Site struct {
    ObjectID   primitive.ObjectID `bson:"_id,omitempty"`
    UserID     primitive.ObjectID `bson:"user_id"`
    SiteID     string             `json:"site_id"     bson:"site_id"`
    IsFavorite bool               `json:"is_favorite" bson:"is_favorite"`
}
