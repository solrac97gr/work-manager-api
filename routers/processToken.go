package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/solrac97gr/yendo/database"
	"github.com/solrac97gr/yendo/models"
)

/*Email : value of the user email*/
var Email string

/*UserID : value of the user id*/
var UserID string

/*ProcessToken : Process token and extract values*/
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("thisisnotwitter")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format not valid")
	}

	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, exist, _ := database.UserExist(claims.Email)
		if exist {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, exist, UserID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}
