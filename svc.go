package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sideshow/apns2/token"
)

const kid = "CUG44HA5T5"

func handler(w http.ResponseWriter, r *http.Request) {

	// TODO don't hit the disk on every req
	authKey, err := token.AuthKeyFromFile("/keybase/private/atec/etc/musickitjs/AuthKey_" + kid + ".p8")
	if err != nil {
		log.Fatal("reading p8 file", err)
	}

	token := &token.Token{
		AuthKey:  authKey,
		IssuedAt: time.Now().Unix(),
		//		"exp":   1546305557,
		KeyID:  kid,
		TeamID: "Y82E2K77P5",
	}

	if ok, _ := token.Generate(); ok {
		log.Println("Bearer", token.Bearer)
		w.Write([]byte(token.Bearer))
	} else {
		log.Println("generating token", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
