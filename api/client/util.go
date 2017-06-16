package client

import (
	"strconv"
	"time"

	"github.com/UnnoTed/hide"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/qwentic/qcrm/api/util"
	"github.com/qwentic/qcrm/config"
)

// CreateToken creates a jwt token with a ID
func CreateToken(hID interface{}, encrypt bool) (string, error) {
	var id string

	switch hID.(type) {
	case hide.Int64:
		id = strconv.FormatInt(int64(hID.(hide.Int64)), 10)
	case string:
		id = hID.(string)
	case int64:
		id = strconv.FormatInt(hID.(int64), 10)
	case uint:
		id = strconv.FormatInt(int64(hID.(uint)), 10)
	}

	var (
		encryptedID = id
		err         error
	)

	if encrypt {
		// encrypt the user's ID
		encryptedID, err = util.Encrypt(id, config.EncryptionKey)
		if err != nil {
			return "", err
		}
	}

	// create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&UserToken{
			encryptedID,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(config.JWTExpirationTime).Unix(),
				Id:        encryptedID,
				IssuedAt:  time.Now().Unix(),
				Issuer:    issuer,
				NotBefore: time.Now().Unix(),
			},
		})

	// sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(config.JWTSecret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// WillTokenExpire checks if a token will expire
// the current range is 5-30 minutes
func WillTokenExpire(expAt int64) bool {
	exp := time.Unix(expAt, 0)
	almostNow := time.Now().Add(5 * time.Minute)
	later := time.Now().Add(30 * time.Minute)

	// after 5min and before 30min
	return exp.After(almostNow) && exp.Before(later)
}

/*// EndTrial changes a user's trial state to false
func EndTrial(uid interface{}) (bool, error) {
	var cl *Client

	switch uid.(type) {
	case *Client:
		cl = uid.(*Client)
	default:

		// get the user
		cl = NewClient()
		res := db.Where("id = ?", uid).First(&cl)
		if res.Error != nil {
			if res.RecordNotFound() {
				return false, errors.New("Client from token doesn't exists")
			}

			return false, res.Error
		}
	}

	// check if the account was created 14 days ago
	created := time.Unix(int64(cl.CreatedAt), 0).Add(14 * 24 * time.Hour)
	if cl.Trial && util.AroundTime(time.Now(), created, 1*time.Hour) {
		cl.Trial = false

		if err := cl.Update(); err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}
*/
