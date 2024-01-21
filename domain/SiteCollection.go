package domain

type SiteCollection struct {
	UserID string `bson:"user_id"`
	Sites  []Site `bson:"sites"`
}
