package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DeleteWork : Get a work by id*/
func DeleteWork(ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("yendo")
	col := db.Collection("works")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": objID}

	res, err := col.DeleteOne(ctx, condition)
	if err != nil {
		return false, err
	}
	if res.DeletedCount == 0 {
		return false, fmt.Errorf("DeleteOne() document not found")
	}
	return true, nil
}
