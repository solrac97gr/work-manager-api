package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Category : This model represent the category type*/
type Category struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	IconType int16              `bson:"icon_type" json:"icon_type"`
}

func (c *Category) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.IconType,validation.Required),
	)
}
