package main

import (
	"delivery.futuroclick.com.br/dao"
	"delivery.futuroclick.com.br/model"
	"delivery.futuroclick.com.br/service"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {

	router := httprouter.New()

	//autenticacao do lojista
	router.GET("/user/:user/password/:password/admin/get", authSeller)

	//autenticacao do cliente
	router.GET("/seller/:id/user/:user/password/:password/consumer/get", authConsumer)

	//recupera as lojas do lojista
	router.GET("/seller/:id/stores/get", getStoresBySeller)

	//salva um pedido do cliente
	router.POST("/seller/:id/consumer/:consumer/orders/post/token/:token", putOrder)

	//recupera os pedidos do cliente
	router.GET("/seller/:id/consumer/:consumer/orders/get", getOrdersByConsumer)

	//recupera os pedidos do lojista
	router.GET("/seller/:id/orders/get", getOrdersBySeller)

	//Cria um novo usuário para o aplicativo
	router.POST("/seller/:id/consumers/post/token/:token", createUser)

	//Cria um novo usuário para o aplicativo
	router.POST("/seller/:id/consumer/:consumer/order/:order/status/:status/update", updateStatusOrder)

	//recupera as lojas do lojista
	router.POST("/seller/:id/time/:time/post/token/:token", putTimeDelivery)

	log.Fatal(http.ListenAndServe(":7002", router))

}

func authSeller(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	user := ps.ByName("user")

	password := ps.ByName("password")

	result := dao.AuthSeller(user, password)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")

	w.Header().Set("Content-Type", "application/json; charset=utf8")

	json.NewEncoder(w).Encode(result)

}

func authConsumer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	user := ps.ByName("user")

	password := ps.ByName("password")

	if err != nil {

		panic(err)

	}

	result := dao.AuthConsumer(int64(id), user, password)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")

	w.Header().Set("Content-Type", "application/json; charset=utf8")

	json.NewEncoder(w).Encode(result)

}

func putOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	token := ps.ByName("token")

	value, err := service.Base64Decode([]byte(token))

	s := string(value[:])

	if err != nil {

		panic(err)

	} else {

		currentTime := time.Now().UnixNano() / int64(time.Millisecond)

		timestamp, err := strconv.Atoi(s)

		if err != nil {

			panic(err)

		}

		if (currentTime+30000) > int64(timestamp) && (currentTime-30000) < int64(timestamp) {

			id, err := strconv.Atoi(ps.ByName("id"))

			entity := model.SellerOrder{}

			json.NewDecoder(r.Body).Decode(&entity)

			result := dao.SaveOrder(int64(id), entity)

			if err != nil {

				panic(err)

			}

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")

			w.Header().Set("Content-Type", "application/json; charset=utf8")

			json.NewEncoder(w).Encode(result)

		} else {

			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

		}

	}

}

func getOrdersByConsumer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	consumer := ps.ByName("consumer")

	if err != nil {

		panic(err)

	}

	result := dao.GetOrdersByConsumer(int64(id), consumer)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")

	w.Header().Set("Content-Type", "application/json; charset=utf8")

	json.NewEncoder(w).Encode(result)

}

func getOrdersBySeller(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	q := r.URL.Query()

	store, err := strconv.Atoi(q.Get("store"))

	order, err := strconv.Atoi(q.Get("order"))

	status, err := strconv.Atoi(q.Get("status"))

	limit, err := strconv.Atoi(q.Get("limit"))

	if err != nil {

		panic(err)

	}

	result := dao.GetOrdersBySeller(int64(id), status, store, int64(order), limit)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")

	w.Header().Set("Content-Type", "application/json; charset=utf8")

	json.NewEncoder(w).Encode(result)

}

func getStoresBySeller(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {

		panic(err)

	}

	result := dao.GetStoresBySeller(int64(id))

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")

	w.Header().Set("Content-Type", "application/json; charset=utf8")

	json.NewEncoder(w).Encode(result)

}

func createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	token := ps.ByName("token")

	value, err := service.Base64Decode([]byte(token))

	s := string(value[:])

	if err != nil {

		panic(err)

	} else {

		current_time := time.Now().UnixNano() / int64(time.Millisecond)

		timestamp, err := strconv.Atoi(s)

		if err != nil {

			panic(err)

		}

		if (current_time+30000) > int64(timestamp) && (current_time-30000) < int64(timestamp) {

			id, err := strconv.Atoi(ps.ByName("id"))

			entity := model.SellerUser{}

			json.NewDecoder(r.Body).Decode(&entity)

			result := dao.SaveUser(int64(id), entity)

			if err != nil {

				panic(err)

			}

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")

			w.Header().Set("Content-Type", "application/json; charset=utf8")

			json.NewEncoder(w).Encode(result)

		} else {

			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

		}

	}

}

func updateStatusOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	token := ps.ByName("token")

	value, err := service.Base64Decode([]byte(token))

	s := string(value[:])

	if err != nil {

		panic(err)

	} else {

		current_time := time.Now().UnixNano() / int64(time.Millisecond)

		timestamp, err := strconv.Atoi(s)

		if err != nil {

			panic(err)

		}

		if (current_time+30000) > int64(timestamp) && (current_time-30000) < int64(timestamp) {

			id, err := strconv.Atoi(ps.ByName("id"))

			consumer := ps.ByName("consumer")

			order, err := strconv.Atoi(ps.ByName("order"))

			status, err := strconv.Atoi(ps.ByName("status"))

			result := dao.UpdateStatusOrder(int64(id), consumer, int64(order), status)

			if err != nil {

				panic(err)

			}

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")

			w.Header().Set("Content-Type", "application/json; charset=utf8")

			json.NewEncoder(w).Encode(result)

		} else {

			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

		}

	}

}

func putTimeDelivery(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	token := ps.ByName("token")

	value, err := service.Base64Decode([]byte(token))

	s := string(value[:])

	if err != nil {

		panic(err)

	} else {

		currentTime := time.Now().UnixNano() / int64(time.Millisecond)

		timestamp, err := strconv.Atoi(s)

		if err != nil {

			panic(err)

		}

		if (currentTime+30000) > int64(timestamp) && (currentTime-30000) < int64(timestamp) {

			id, err := strconv.Atoi(ps.ByName("id"))

			time, err := strconv.Atoi(ps.ByName("time"))

			if err != nil {

				panic(err)

			}

			entity := model.SellerOrder{}

			json.NewDecoder(r.Body).Decode(&entity)

			result := dao.SaveTimeDelivery(int64(id), time)

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")

			w.Header().Set("Content-Type", "application/json; charset=utf8")

			json.NewEncoder(w).Encode(result)

		} else {

			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

		}

	}

}
