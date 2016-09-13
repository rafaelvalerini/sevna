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

	router.GET("/v1/user/auth", auth)

	router.GET("/v1/user/reset/password", resetPassword)

	router.POST("/v1/user/create", createUser)

	router.GET("/v1/estimates/count", countEstimates)

	router.GET("/v1/estimates/modality/selected/count", countModalities)

	router.GET("/v1/notification", notification)

	router.GET("/v1/estimates/saved", savedEstimates)

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

func auth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	q := r.URL.Query()

	mail := q.Get("mail")

	password := q.Get("password")

	if auth != "65edc9b5-d134-4c8b-9be5-ee2c722f4a54" {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if password == "" || mail == ""{

		http.Error(w, "User or Password not found", http.StatusBadRequest)

	}else{

		result := repository.Login(mail, password);

		if result.Id <= 0 {
			
			http.Error(w, "User or Password not found", http.StatusNotFound)

		}else{

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(result)

		}

	}

}

func resetPassword(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	q := r.URL.Query()

	mail := q.Get("mail")

	if auth != "65edc9b5-d134-4c8b-9be5-ee2c722f4a54" {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if mail == ""{

		http.Error(w, "Mail not found", http.StatusBadRequest)

	}else{

		result := service.ResetPassword(mail);

		if result.Id <= 0 {
			
			http.Error(w, "User or Password not found", http.StatusNotFound)

		}else{

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(result)

		}

	}

}

func createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	entity := model.User{}

	json.NewDecoder(r.Body).Decode(&entity)

	if auth != "65edc9b5-d134-4c8b-9be5-ee2c722f4a54" {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if entity.Mail == "" || entity.Password == ""{

		http.Error(w, "User or Password not found", http.StatusBadRequest)

	}else{

		result := repository.FindUserByMail(entity.Mail);

		if result.Id > 0{

			http.Error(w, "User already registered", http.StatusBadRequest)

		}else{

			repository.CreateUser(entity)

			userFinal := repository.FindUserByMail(entity.Mail);

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(userFinal)

		}

	}

}


func countEstimates(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	if auth != "65edc9b5-d134-4c8b-9be5-ee2c722f4a54" {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else{

		result := repository.CountEstimates();

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(model.MetaInt64{Value: result})

	}

}

func countModalities(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	if auth != "65edc9b5-d134-4c8b-9be5-ee2c722f4a54" {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else{

		result := repository.CountModalities();

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(result)

	}

}

func notification(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	entity := model.Message{}

	json.NewDecoder(r.Body).Decode(&entity)

	if auth != "65edc9b5-d134-4c8b-9be5-ee2c722f4a54" {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if entity.Message == "" {
		
		http.Error(w, "Message not found", http.StatusBadRequest)

	}else{

		q := r.URL.Query()

		state := q.Get("state")

		city := q.Get("city")

		result := service.SendNotification(state, city, entity.Message);

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(result)

	}

}

func savedEstimates(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	if auth != "65edc9b5-d134-4c8b-9be5-ee2c722f4a54" {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else{

		result := repository.SumSavedEstimates();

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(result)

	}

}