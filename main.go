package main

import (
	"agregador/service"
	"agregador/model"
	"agregador/repository"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {

	router := httprouter.New()
	
	router.POST("/v1/estimates", estimate)

	router.POST("/v1/estimate/selected/:selected", selected)

	c := cors.New(cors.Options{
	    AllowedOrigins: []string{"*"},
	    AllowCredentials: true,
	    AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))

}

func estimate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	entity := model.RequestAggregator{}

	json.NewDecoder(r.Body).Decode(&entity)

	if auth != "65edc9b5-d134-4c8b-9be5-ee2c722f4a54"{

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if entity.Start.Lat == 0 {

		http.Error(w, "Start Latitude not found", http.StatusBadRequest)

	}else if entity.Start.Lng == 0 {

		http.Error(w, "Start Longitude not found", http.StatusBadRequest)

	}else if entity.End.Lat == 0 {

		http.Error(w, "End Latitude not found", http.StatusBadRequest)

	}else if entity.End.Lng == 0 {

		http.Error(w, "End Longitude not found", http.StatusBadRequest)

	}else{

		result := service.AgregateAll(entity)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(result)

	}

}

func selected(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	if auth != "65edc9b5-d134-4c8b-9be5-ee2c722f4a54" {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if ps.ByName("selected") == "" {

		http.Error(w, "Selected not found", http.StatusBadRequest)

	}else{

		result := repository.Selected(ps.ByName("selected"));

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(result)

	}

}