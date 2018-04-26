package main

import (

	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/remind101/newrelic"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {

	newrelic.Init("VAH-GOLANG", "d8fd304dc5c8d8ddab8ee4471263c21ccd989ceb")
	tx := newrelic.NewTx("/v1/estimate")
	tx.Start()
	defer tx.End()

	router := httprouter.New()

	router.GET("/v1/ping", pong)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))

}

func pong(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("pong")

}

