package main

import (
	"crypto/ecdsa"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/sideshow/apns2/token"
)

const (
	etc = "/keybase/private/atec/etc"
	cat = "kitjs/AuthKey_"
	ext = ".p8"

	alg = "ES256"
	iss = "Y82E2K77P5"
)

var (
	kids = map[string]string{
		"music": "CUG44HA5T5",
		"map":   "YKVC29UG5H",
	}
	keys = make(map[string]*ecdsa.PrivateKey)
)

func init() {
	for svc, kid := range kids {

		path := etc + "/" + svc + cat + kid + ext
		log.Println("loading auth key from path", path)

		key, err := token.AuthKeyFromFile(path)
		if err != nil {
			log.Fatal("reading p8 file", err)
		}

		keys[svc] = key
	}
}

func sign(w http.ResponseWriter, r *http.Request) {

	svc := mux.Vars(r)["svc"]
	log.Println("svc", svc)

	if svc != "music" && svc != "map" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jwtToken := &jwt.Token{
		Header: map[string]interface{}{
			"alg": alg,
			"kid": kids[svc],
		},
		Claims: jwt.MapClaims{
			"iss": iss,
			"exp": time.Now().Unix() + 3000,
		},
		Method: jwt.SigningMethodES256,
	}

	bearer, err := jwtToken.SignedString(keys[svc])
	if err != nil {
		log.Println("signing bearer", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(bearer))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{svc}", sign)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
