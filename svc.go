package main

import (
	"log"

	"github.com/sideshow/apns2/token"
)

func main() {

	authKey, err := token.AuthKeyFromFile("/keybase/private/atec/etc/musickitjs/AuthKey_CUG44HA5T5.p8")
	if err != nil {
		log.Fatal("reading p8 file", err)
	}

	token := &token.Token{
		AuthKey: authKey,
		//		"exp":   1546305557,
		KeyID:  "YKVC29UG5H",
		TeamID: "Y82E2K77P5",
	}

	if ok, _ := token.Generate(); ok {
		println("Bearer", token.Bearer)
	} else {
		log.Println("generating token", err)
	}
}
