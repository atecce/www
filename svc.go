package main

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sideshow/apns2/token"
)

const kid = "CUG44HA5T5"

func handler(w http.ResponseWriter, r *http.Request) {

	// TODO don't hit the disk on every req
	authKey, err := token.AuthKeyFromFile("/keybase/private/atec/etc/musickitjs/AuthKey_" + kid + ".p8")
	if err != nil {
		log.Fatal("reading p8 file", err)
	}

	jwtToken := &jwt.Token{
		Header: map[string]interface{}{
			"alg": "ES256",
			"kid": kid,
		},
		Claims: jwt.MapClaims{
			"iss": "Y82E2K77P5",
			"exp": time.Now().Unix() + 3000,
		},
		Method: jwt.SigningMethodES256,
	}
	bearer, _ := jwtToken.SignedString(authKey)
	log.Println(bearer)
	w.Write([]byte(bearer))
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
