package main

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sideshow/apns2/token"
)

const (
	etc = "/keybase/private/atec/etc"
	cat = "kitjs/AuthKey_"
	ext = ".p8"

	alg = "ES256"
	iss = "Y82E2K77P5"
)

var kids = map[string]string{
	"music": "CUG44HA5T5",
	"map":   "YKVC29UG5H",
}

func sign(temp string) string {

	kid := kids[temp]
	log.Println("kid", kid)

	// TODO don't hit the disk on every req
	path := etc + "/" + temp + cat + kid + ext
	log.Println("loading auth key from path", path)
	authKey, err := token.AuthKeyFromFile(path)
	if err != nil {
		log.Fatal("reading p8 file", err)
	}

	jwtToken := &jwt.Token{
		Header: map[string]interface{}{
			"alg": alg,
			"kid": kid,
		},
		Claims: jwt.MapClaims{
			"iss": iss,
			"exp": time.Now().Unix() + 3000,
		},
		Method: jwt.SigningMethodES256,
	}
	bearer, _ := jwtToken.SignedString(authKey)
	return bearer
}

func mapkit(w http.ResponseWriter, r *http.Request) {

	bearer := sign("map")
	log.Println(bearer)
	w.Write([]byte(bearer))
}

func musickit(w http.ResponseWriter, r *http.Request) {

	bearer := sign("music")
	log.Println(bearer)
	w.Write([]byte(bearer))
}

func main() {

	http.HandleFunc("/map", mapkit)
	http.HandleFunc("/music", musickit)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
