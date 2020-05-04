package database

import (
	"context"
	"errors"
	"time"

	"github.com/solrac97gr/yendoapi/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*UpdateWork : Update a existe work*/
func UpdateWork(work models.Work) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("yendo")
	col := db.Collection("works")

	register := make(map[string]interface{})

	if len(work.Name) > 0 {
		register["name"] = work.Name
	}
	if len(work.Description) > 0 {
		register["description"] = work.Description
	}

	register["create_at"] = work.CreateAt

	if len(work.AddressID) > 0 {
		register["address_id"] = work.AddressID
	}
	if len(work.CategoryID) > 0 {
		register["category_id"] = work.CategoryID
	}

	updateString := bson.M{
		"$set": register,
	}

	objtID := work.ID

	filter := bson.M{"_id": bson.M{"$eq": objtID}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, errors.New("Not updated")
	}
	return true, nil
}
