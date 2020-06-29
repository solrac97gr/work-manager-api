package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

/*SearchWork : Get all work by user id*/
func GetAll(ID string) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	var results []bson.M
	db := MongoCN.Database("yendo")
	col := db.Collection("works")


	cursor, err := col.Find(ctx, bson.D{{"user_id", ID}})
	if err != nil {
		fmt.Println("si es aca")
		fmt.Println("Registros no encontrados " + err.Error())
		return results, err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println("Registros no se pudieron obtener " + err.Error())
		return results, err
	}
	return results, nil
}
