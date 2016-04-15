package main

import (
	"cevafacil.com.br/dao"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	//log "github.com/dmuth/google-go-log4go"
)

func main() {

	//log.SetLevel(log.DebugLevel)

	//log.SetDisplayTime(true)

	//log.Info("cevafacil")

	router := httprouter.New()

	router.GET("/type/:id/brand/:brand/size/:size/lat/:lat/lng/:lng", findByIdAndSize)

	router.GET("/seller/:id/:size/:lat/:lng", findBySeller)

	router.GET("/user/:user/password/:pass", findByUserAndPassword)

	log.Fatal(http.ListenAndServe(":7001", router))

}

func findByIdAndSize(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	brand, err := strconv.Atoi(ps.ByName("brand"))

	size, err := strconv.Atoi(ps.ByName("size"))

	lat, err := strconv.ParseFloat(ps.ByName("lat"), 64)

	lng, err := strconv.ParseFloat(ps.ByName("lng"), 64)

	if err != nil {

		panic(err)

	}

	result := dao.FindSellerByIdAndSize(id, brand, size, lat, lng)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json; charset=utf8")

	json.NewEncoder(w).Encode(result)

}

func findBySeller(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	size, err := strconv.Atoi(ps.ByName("size"))

	lat, err := strconv.ParseFloat(ps.ByName("lat"), 64)

	lng, err := strconv.ParseFloat(ps.ByName("lng"), 64)

	if err != nil {

		panic(err)

	}

	result := dao.FindSellerById(id, size, lat, lng)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json; charset=utf8")

	json.NewEncoder(w).Encode(result)

}

func findByUserAndPassword(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	user := ps.ByName("user")

	password := ps.ByName("pass")

	result := dao.FindSellerByUserAndPassword(user, password)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json; charset=utf8")

	json.NewEncoder(w).Encode(result)

}
