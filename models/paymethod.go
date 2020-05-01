package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*PayMethod : The model represent the pay method*/
type PayMethod struct {
	ID          primitive.ObjectID
	Name        string
	Description string
	URL         string
}
