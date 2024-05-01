package auth

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenrateJWt(email string) (map[string]string, error) {
	expirationtime := time.Now().Add(1 * time.Hour)
	Claims := &JWTClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationtime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenstring, err := token.SignedString([]byte(os.Getenv("SUPERSECRET")))
	if err != nil {
		return nil, err
	}
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["email"] = email
	rtClaims["exp"] = time.Now().Add(24 * time.Hour)

	rt, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  tokenstring,
		"Refresh_tiken": rt,
	}, nil

}

func Validtoken(signedToken string) (err error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SUPERSECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Local().Unix() {
			err = errors.New("token expired")
			return
		}

	}
	return

}
