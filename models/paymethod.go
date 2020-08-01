package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*PayMethod : The model represent the pay method*/
type PayMethod struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	URL         string             `bson:"url" json:"url,omitempty"`
}

func (pm *PayMethod) Validate() error{
	return validation.ValidateStruct(pm,
		validation.Field(&pm.Name, validation.Required),
		validation.Field(&pm.Description,validation.Required),
		validation.Field(&pm.URL,validation.Required,is.URL),
	)
}