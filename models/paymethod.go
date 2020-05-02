package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*PayMethod : The model represent the pay method*/
type PayMethod struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	URL         string             `bson:"url" json:"url,omitempty"`
}
