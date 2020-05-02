package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Work : The model represent the job*/
type Work struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	Price       float64            `bson:"price" json:"price,omitempty"`
	Category    Category           `bson:"category" json:"category,omitempty"`
	CreateAt    time.Time          `bson:"create_at" json:"create_at,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
	AddressID   string             `bson:"address_id" json:"address_id,omitempty"`
}
