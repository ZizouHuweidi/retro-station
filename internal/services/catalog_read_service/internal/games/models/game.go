package models

import (
	"time"
)

type Game struct {
	// we generate id ourselves because auto generate mongo string id column with type _id is not an uuid
	Id          string    `json:"id"                    bson:"_id,omitempty"` // https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/write-operations/insert/#the-_id-field
	GameId      string    `json:"gameId"             bson:"gameId"`
	Name        string    `json:"name,omitempty"        bson:"name,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Price       float64   `json:"price,omitempty"       bson:"price,omitempty"`
	Genre       string    `json:"genre,omitempty"       bson:"genre,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"   bson:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"   bson:"updatedAt,omitempty"`
}

type GamesList struct {
	TotalCount int64   `json:"totalCount" bson:"totalCount"`
	TotalPages int64   `json:"totalPages" bson:"totalPages"`
	Page       int64   `json:"page"       bson:"page"`
	Size       int64   `json:"size"       bson:"size"`
	Games      []*Game `json:"games"   bson:"games"`
}
