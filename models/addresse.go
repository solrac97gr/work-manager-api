package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Address : The model represent the job address*/
type Address struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	District    string             `bson:"district" json:"district,omitempty"`
	FullAddress string             `bson:"full_address" json:"full_address,omitempty"`
	Reference   string             `bson:"reference" json:"reference,omitempty"`
	CreateAt    time.Time          `bson:"create_at" json:"create_at,omitempty"`
	Date        time.Time          `bson:"date" json:"date,omitempty"`
	WorkID      string             `bson:"work_id" json:"work_id,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
}

func (a *Address) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.District, validation.Required),
		validation.Field(&a.FullAddress, validation.Required),
		validation.Field(&a.Reference, validation.Required, validation.Min(20)),
		validation.Field(&a.CreateAt, validation.Required),
		validation.Field(&a.Date, validation.Required),
		validation.Field(&a.WorkID, validation.Required),
		validation.Field(&a.UserID, validation.Required),
	)
}
