package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"log/syslog"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/sideshow/apns2/token"
)

const (
	cat = "kitjs/AuthKey_"
	ext = ".p8"

	alg = "ES256"
	iss = "Y82E2K77P5"
)

var (
	logger *syslog.Writer

	etc = os.Getenv("ETC")

	kids = map[string]string{
		"music": "CUG44HA5T5",
		"map":   "YKVC29UG5H",
	}
	keys = make(map[string]*ecdsa.PrivateKey)
)

func checkSyslog(err error) {
	if err != nil {
		fmt.Println("failed to write to syslog")
	}
}

func init() {

	var err error

	logger, err = syslog.Dial("tcp", "35.231.223.50:514", syslog.LOG_INFO, "auth")
	if err != nil {
		log.Fatal("dialing syslog", err)
	}

	for svc, kid := range kids {

		path := etc + "/" + svc + cat + kid + ext
		checkSyslog(logger.Info("loading auth key from path: " + path))

		key, err := token.AuthKeyFromFile(path)
		if err != nil {
			checkSyslog(logger.Err("reading p8 file: " + err.Error()))
		}

		keys[svc] = key
	}
}

func middleware(w http.ResponseWriter, r *http.Request) bool {

	checkSyslog(logger.Info(r.Method + " " + r.URL.Path + " " + r.RemoteAddr + " " + r.Host))

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	return true
}

func hit(w http.ResponseWriter, r *http.Request) {

	if ok := middleware(w, r); !ok {
		return
	}

	w.WriteHeader(http.StatusOK)
}

func sign(w http.ResponseWriter, r *http.Request) {

	if ok := middleware(w, r); !ok {
		return
	}

	svc := mux.Vars(r)["svc"]
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
		checkSyslog(logger.Err("signing bearer: " + err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Write([]byte(bearer))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusMovedPermanently)
}

func main() {

	go func() {
		checkSyslog(logger.Err(http.ListenAndServe(":80", http.HandlerFunc(redirect)).Error()))
	}()

	r := mux.NewRouter()
	r.HandleFunc("/", hit)
	r.HandleFunc("/{svc}", sign)

	authEtc := etc + "auth"

	checkSyslog(logger.Err(http.ListenAndServeTLS(":443", authEtc+"/server.crt", authEtc+"/server.key", r).Error()))
}
