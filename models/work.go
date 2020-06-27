package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Work : The model represent the job*/
type Work struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	Price       float64            `bson:"price" json:"price,omitempty"`
	CreateAt    time.Time          `bson:"create_at" json:"create_at,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
	AddressID   string             `bson:"address_id" json:"address_id,omitempty"`
	CategoryID  string             `bson:"category_id" json:"category_id,omitempty"`
}

func (w Work) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.Name, validation.Required,validation.Length(5, 50)),
		validation.Field(&w.Description, validation.Required,validation.Length(20, 50)),
		validation.Field(&w.Price, validation.Required, validation.Min(20.00),validation.Max(400.00)),
		validation.Field(&w.CreateAt,validation.Required),
		validation.Field(&w.UserID, validation.Required),
		validation.Field(&w.AddressID, validation.Required),
		validation.Field(&w.CategoryID, validation.Required),
	)
}