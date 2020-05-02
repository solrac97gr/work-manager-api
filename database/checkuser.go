package database

import (
	"context"
	"time"

	"github.com/solrac97gr/yendoapi/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*UserExist : Check if the user is alredy register*/
func UserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)

	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
