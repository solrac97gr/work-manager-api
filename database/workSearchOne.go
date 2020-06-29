package database

import (
	"context"
	"fmt"
	"time"

	"github.com/solrac97gr/yendoapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*SearchWork : Get a work by id*/
func SearchWork(ID string) (models.Work, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("yendo")
	col := db.Collection("works")

	var work models.Work
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": objID}

	err := col.FindOne(ctx, condition).Decode(&work)

	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return work, err
	}
	return work, nil
}
