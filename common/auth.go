package common

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
)

var (
	verifyKey, signKey []byte
)

// AppClaims provides custom claim for JWT
type AppClaims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

//Read the key files before starting http handlers

func initKeys() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	log.Println(signKey)
	if err != nil {
		log.Fatalf("[initkeys]: %s\n", err)
	}
	verifyKey, err := ioutil.ReadFile(pubKeyPath)
	log.Println(verifyKey)
	if err != nil {
		log.Fatalf("[verifyKey]: %s\n", err)
	}
}

//generate JWT token

func GenerateJWT(name, role string) (string, error) {
	//create a signer for rsa 256

	//t := jwt.New(jwt.GetSigningMethod("RS256"))
	//set claims for JWT token

	//t.Claims["iss"] = "admin"
	t := jwt.New(jwt.SigningMethodRS512)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	claims["iat"] = time.Now().Unix()
	t.Claims = claims
	//t.Claims["UserInfo"] = struct {
	//	Name string
	//	Role string
	//}{name, role}

	// set the expire time for JWT token
	//t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//Middleware for validating JWT token

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//validate the token

	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		//verify the token with public key, which is the counter part of the private key
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError: //JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				DisplayAppError(
					w,
					err,
					"Access token is expired",
					401,
				)
				return
			default:
				DisplayAppError(w,
					err,
					"Error while parsing the token",
					500,
				)
				return
			}
		default:
			DisplayAppError(w,
				err,
				"Error while parsing Access Token",
				500)
			return
		}
	}
	if token.Valid {
		next(w, r)
	} else {
		DisplayAppError(w,
			err,
			"invalid Access Token",
			401,
		)
	}
}
