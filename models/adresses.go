package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Addresse : The model represent the job adress*/
type Addresse struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Distrit     string             `bson:"district" json:"distric,omitempty"`
	FullAddress string             `bson:"full_address" json:"full_address,omitempty"`
	Reference   string             `bson:"reference" json:"reference,omitempty"`
	CreateAt    time.Time          `bson:"create_at" json:"create_at,omitempty"`
	Date        time.Time          `bson:"date" json:"date,omitempty"`
	WorkID      string             `bson:"work_id" json:"work_id,omitempty"`
	UserID      string             `bson:"user_id" json:"user_id,omitempty"`
}
