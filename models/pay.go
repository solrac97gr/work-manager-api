package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Pay : The model pay represent the user payment object*/
type Pay struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Amount      float32            `bson:"amount" json:"amount,omitempty"`
	Date        time.Time          `bson:"date" json:"date,omitempty"`
	PayMethodID string             `bson:"pay_method_id" json:"pay_method_id,omitempty"`
	WorkID      string             `bson:"work_id" json:"work_id,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
}

func (p *Pay) Validate() error {
	return validation.ValidateStruct(p,
		validation.Field(&p.Amount, validation.Min(20)),
		validation.Field(&p.Date,validation.Required),
		validation.Field(&p.PayMethodID,validation.Required),
		validation.Field(&p.WorkID,validation.Required),
		validation.Field(&p.UserID,validation.Required),
	)
}
