package domain

type Site struct {
    SiteID string   `json:"site_id"     bson:"site_id"`
    IsFavorite bool `json:"is_favorite" bson:"is_favorite"`
    Paths []string  `json:"paths"       bson:"paths"`
}
