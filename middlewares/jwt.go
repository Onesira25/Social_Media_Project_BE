package middlewares

import (
	"Social_Media_Project_BE/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id string) (string, error) {
	var data = jwt.MapClaims{}
	data["id"] = id
	// data["username"] = username
	data["iat"] = time.Now().Unix()
	data["exp"] = time.Now().Add(time.Hour * 3).Unix()

	var proccessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	result, err := proccessToken.SignedString([]byte(config.JWTSECRET))
	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}

func DecodeToken(i interface{}) string {
	var claim = i.(*jwt.Token).Claims.(jwt.MapClaims)
	var result string

	// if val, found := claim["username"]; found {}

	if val, found := claim["id"]; found {
		result = val.(string)
	}

	return result
}
