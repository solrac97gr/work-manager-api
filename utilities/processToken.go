package utilities

import (
	"errors"
	database "github.com/solrac97gr/yendoapi/database/user"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/solrac97gr/yendoapi/models"
)

/*Email : value of the user email*/
var Email string

/*UserID : value of the user id*/
var UserID string

/*ProcessToken : Process jwt and extract values*/
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("thisisnotyendo")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("jwt format not valid")
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
		return claims, false, string(""), errors.New("invalid jwt")
	}
	return claims, false, string(""), err
}
