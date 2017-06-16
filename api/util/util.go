package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/UnnoTed/hide"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fatih/structs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/qwentic/qcrm/config"
)

// KeySize ya
const KeySize = 16

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomCIV generates a random common iv
func RandomCIV() ([]byte, error) {
	civ := make([]byte, KeySize)
	_, err := rand.Read(civ)
	return civ, err
}

// Encrypt .
func Encrypt(text string, key []byte) (string, error) {
	text = hex.EncodeToString([]byte(text))
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	civ, err := RandomCIV()
	if err != nil {
		return "", err
	}

	e := cipher.NewCFBEncrypter(c, civ)
	enc := make([]byte, len(text))
	e.XORKeyStream(enc, []byte(text))

	text = hex.EncodeToString(append(civ, enc...))
	return text, nil
}

// Decrypt .
func Decrypt(text string, key []byte) (string, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	block := c.BlockSize() * 2

	civ, err := hex.DecodeString(text[:block])
	if err != nil {
		return "", err
	}

	t, err := hex.DecodeString(text[block:])
	if err != nil {
		return "", err
	}
	text = string(t)

	d := cipher.NewCFBDecrypter(c, []byte(civ))
	dec := make([]byte, len(text))
	d.XORKeyStream(dec, []byte(text))

	t, err = hex.DecodeString(string(dec))
	if err != nil {
		return "", err
	}
	return string(t), nil
}

// GetUserID decrypts the user id from the user token stored in echo's context
func GetUserID(c echo.Context) (int64, error) {
	usr := c.Get(middleware.DefaultJWTConfig.ContextKey).(*jwt.Token).Claims //.(*client.UserToken)
	m := structs.Map(usr)

	id, err := Decrypt(m["UID"].(string), config.EncryptionKey)
	if err != nil {
		return 0, err
	}

	log.Println("--------ID:", m["UID"])
	return strconv.ParseInt(id, 10, 64)
}

func Obfuscate(id uint) int64 {
	return hide.Int64Obfuscate(int64(id), nil, nil)
}
