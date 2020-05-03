package database

import (
	"context"
	"time"

	"github.com/solrac97gr/yendoapi/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*RegisterWork : Register a work in the database*/
func RegisterWork(work models.Work) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("yendo")
	col := db.Collection("works")

	result, err := col.InsertOne(ctx, work)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
