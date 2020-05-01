package models

import (
	"time"

	"github.com/solrac97gr/yendo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Pay : The model pay represent the user payment object*/
type Pay struct {
	ID        primitive.ObjectID
	PayMethod models.PayMethod
	Amount    float32
	Date      time.Time
	WorkID    primitive.ObjectID
	UserID    primitive.ObjectID
}
