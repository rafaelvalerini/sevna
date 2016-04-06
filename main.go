package main

import (
	"cevafacil.com.br/dao"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func main() {

	router := httprouter.New()

	router.GET("/types/:id/:size/:lat/:lng", findByIdAndSize)

	log.Fatal(http.ListenAndServe(":7001", router))

}

func findByIdAndSize(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	size, err := strconv.Atoi(ps.ByName("size"))

	lat, err := strconv.ParseFloat(ps.ByName("lat"), 64)

	lng, err := strconv.ParseFloat(ps.ByName("lng"), 64)

	if err != nil {

		panic(err)

	}

	result := dao.FindSellerByIdAndSize(id, size, lat, lng)

	w.Header().Set("Content-Type", "application/json; charset=utf8")

	json.NewEncoder(w).Encode(result)

}
