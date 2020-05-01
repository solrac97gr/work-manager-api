package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Adresse : The model represent the job adress*/
type Adresse struct {
	ID         primitive.ObjectID
	Distrit    string
	FullAdress string
	Reference  string
	CreateAt   time.Time
	Date       time.Time
	WorkID     primitive.ObjectID
	UserID     primitive.ObjectID
}
