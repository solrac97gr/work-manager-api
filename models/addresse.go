package models

import (
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
