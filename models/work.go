package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Work : The model represent the job*/
type Work struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Description string             `bson:"description" json:"description,omitempty"`
	PhotoUrl    string             `bson:"photo_url" json:"photo_url,omitempty"`
	CreateAt    time.Time          `bson:"create_at" json:"create_at,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
	AddressID   string             `bson:"address_id" json:"address_id,omitempty"`
	CategoryID  string             `bson:"category_id" json:"category_id,omitempty"`
	PayMethodID string             `bson:"pay_method_id" json:"pay_method_id,omitempty"`
}

func (w Work) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.Description, validation.Required, validation.Length(20, 50)),
		validation.Field(&w.PhotoUrl,validation.Required,is.URL),
		validation.Field(&w.CreateAt, validation.Required),
		validation.Field(&w.UserID, validation.Required),
		validation.Field(&w.AddressID, validation.Required),
		validation.Field(&w.CategoryID, validation.Required),
		validation.Field(&w.PayMethodID,validation.Required),
	)
}
