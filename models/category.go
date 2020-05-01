package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*Category : This model represent the category type*/
type Category struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	IconType int16              `bson:"icon_type" json:"icon_type"`
}
