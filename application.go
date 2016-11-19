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
	"strconv"
	"github.com/remind101/newrelic"
)

func main() {

	newrelic.Init("VAH-GOLANG", "d8fd304dc5c8d8ddab8ee4471263c21ccd989ceb")
    tx := newrelic.NewTx("/v1/estima")
    tx.Start()
    defer tx.End()

	router := httprouter.New()

	router.GET("/v1/ping", pong)
	
	router.POST("/v1/estimates", estimate)

	router.POST("/v1/estimate/selected/:selected", selected)

	router.GET("/v1/user/auth", auth)

	router.GET("/v1/user/reset/password", resetPassword)

	router.POST("/v1/user/create", createUser)

	router.GET("/v1/estimates/count", countEstimates)

	router.GET("/v1/estimates/modality/selected/count", countModalities)

	router.GET("/v1/notification", notification)

	router.GET("/v1/estimates/saved", savedEstimates)

	router.POST("/v1/player", savePlayer)

	router.GET("/v1/players", findPlayers)

	router.POST("/v1/player/:player/delete", deletePlayer)

	router.POST("/v1/player/:player/modality/:modality/delete", deleteModality)

	router.POST("/v1/player/:player/modality/", saveModality)

	router.POST("/v1/player/:player/modality/:modality/promotion", savePromotion)

	router.POST("/v1/promotion/:promotion/delete", deletePromotion)

	router.GET("/v1/player/:player/modality/:modality/promotions", findPromotion)

	router.GET("/v1/states", getAllStates)

	router.GET("/v1/state/:state/cities", getCityByState)

	router.GET("/v1/estimates/analytics", getAnalytics)

	router.GET("/v1/estimates/promotions/count", countPromotions)

	c := cors.New(cors.Options{
	    AllowedOrigins: []string{"*"},
	    AllowCredentials: true,
	    AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))

}

func pong(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("pong")

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

		q := r.URL.Query()

		promotion := q.Get("promotion")

		store := q.Get("store")

		result := repository.Selected(ps.ByName("selected"), promotion, store);

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

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else{

		result := repository.CountEstimates();

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(model.MetaInt64{Value: result})

	}

}

func countModalities(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

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

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

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

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else{

		result := repository.SumSavedEstimates();

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(result)

	}

}

func savePlayer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	entity := model.Player{}

	json.NewDecoder(r.Body).Decode(&entity)

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if entity.Name == ""{

		http.Error(w, "Name not found", http.StatusBadRequest)

	}else{

		repository.SavePlayer(entity);

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(entity)

	}

}

func deletePlayer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	player, err := strconv.Atoi(ps.ByName("player"))

	if err != nil {

		http.Error(w, "Player not found", http.StatusBadRequest)

	}

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if player <= 0{

		http.Error(w, "Player not found", http.StatusBadRequest)

	}else{

		repository.DeletePlayer(int64(player));

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(model.Meta{Value: "OK"})

	}

}

func deleteModality(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	player, err := strconv.Atoi(ps.ByName("player"))

	if err != nil {

		http.Error(w, "Player not found", http.StatusBadRequest)

	}

	modality, err := strconv.Atoi(ps.ByName("modality"))

	if err != nil {

		http.Error(w, "Modality not found", http.StatusBadRequest)

	}

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if player <= 0 || modality <= 0{

		http.Error(w, "Player not found", http.StatusBadRequest)

	}else{

		repository.DeleteModality(int64(player), int64(modality));

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(model.Meta{Value: "OK"})

	}

}

func findPlayers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else{

		result := repository.FindAllPlayers();

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(result)

	}

}

func savePromotion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	entity := model.Promotion{}

	json.NewDecoder(r.Body).Decode(&entity)

	player, err := strconv.Atoi(ps.ByName("player"))

	if err != nil {

		http.Error(w, "Player not found", http.StatusBadRequest)

	}

	modality, err := strconv.Atoi(ps.ByName("modality"))

	if err != nil {

		http.Error(w, "Modality not found", http.StatusBadRequest)

	}

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if entity.Name == ""{

		http.Error(w, "Name not found", http.StatusBadRequest)

	}else if player <= 0 || modality <= 0{

		http.Error(w, "Player and Modality not found", http.StatusBadRequest)

	}else{

		repository.SavePromotion(entity, modality);

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(entity)

	}

}

func deletePromotion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	promotion, err := strconv.Atoi(ps.ByName("promotion"))

	if err != nil {

		http.Error(w, "promotion not found", http.StatusBadRequest)

	}

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if promotion <= 0{

		http.Error(w, "Promotion not found", http.StatusBadRequest)

	}else{

		repository.DeletePromotion(promotion);

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(model.Meta{Value: "OK"})

	}

}

func findPromotion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	player, err := strconv.Atoi(ps.ByName("player"))

	if err != nil {

		http.Error(w, "Player not found", http.StatusBadRequest)

	}

	modality, err := strconv.Atoi(ps.ByName("modality"))

	if err != nil {

		http.Error(w, "Modality not found", http.StatusBadRequest)

	}

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if player <= 0 || modality <= 0{

		http.Error(w, "Player and Modality not found", http.StatusBadRequest)

	}else{

		result := repository.FindPromotion(player, modality);

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(result)

	}

}

func getAllStates(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else{

		result := repository.FindAllStates();

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(result)

	}

}

func getCityByState(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	state, err := strconv.Atoi(ps.ByName("state"))

	if err != nil {

		http.Error(w, "State not found", http.StatusBadRequest)

	}

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if state <= 0{

		http.Error(w, "State not found", http.StatusBadRequest)

	}else{

		result := repository.FindCityByState(state)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(result)

	}

}

func getAnalytics(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	user := repository.FindUserByToken(auth)

	q := r.URL.Query()

	startAt := q.Get("startAt")

	endAt := q.Get("endAt")

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if startAt == "" || endAt == ""{

		http.Error(w, "Start Date and End Date not found", http.StatusBadRequest)

	}else{

		city := q.Get("city")

		state := q.Get("state")

		player := q.Get("player")

		modality := q.Get("modality")

		result := repository.FindAnalytics(state, city, player, modality, startAt, endAt)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(result)

	}

}

func saveModality(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	entity := model.Modality{}

	json.NewDecoder(r.Body).Decode(&entity)

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else if entity.Name == ""{

		http.Error(w, "Name not found", http.StatusBadRequest)

	}else{

		repository.SaveModality(entity);

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(entity)

	}

}

func countPromotions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	auth := r.Header.Get("Authorization");

	user := repository.FindUserByToken(auth)

	if user.Id <= 0 {

		http.Error(w, "Auth failed", http.StatusUnauthorized)

	}else{

		result := repository.CountPromotions();

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(result)

	}

}