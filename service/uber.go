package service

import (
	"agregador/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

const (
	UBER_DOMAIN             = "https://api.uber.com"
	UBER_URL_PRODUCTS       = "/v1/products"
	UBER_URL_ESTIMATE_TIME  = "/v1/estimates/time"
	UBER_URL_ESTIMATE_PRICE = "/v1/estimates/price"
	UBER_HEADER_AUTH        = "Authorization"
)

func GetEstimatesUber(start_lat float64, start_lng float64, end_lat float64, end_lng float64, player model.Player, token model.TokenPlayer) (response []model.Player) {

	if player.Active == 1 {

		runtime.GOMAXPROCS(3)

		var wg sync.WaitGroup

		wg.Add(3)

		var products model.ResponseProduct

		var times model.ResponseTime

		var prices model.ResponsePrices

		go func() {
			defer wg.Done()

			products = getProducts(start_lat, start_lng, player, token)

		}()

		go func() {
			defer wg.Done()

			times = getTimes(start_lat, start_lng, player, token)

		}()

		go func() {
			defer wg.Done()

			prices = getPrices(start_lat, start_lng, end_lat, end_lng, player, token)

		}()

		wg.Wait()

		return processEstimates(products, times, prices, player)

	} else {

		return response

	}

}

func processEstimates(products model.ResponseProduct, times model.ResponseTime, prices model.ResponsePrices, player model.Player) (response []model.Player) {

	if len(products.Products) > 0 {

		for _, product := range products.Products {

			modal := GetModalityByName(player.Modalities, product.DisplayName, "")

			if modal.Name == "" || modal.Active == 0 {

				continue

			}

			m := model.Player{}

			uuid, err := exec.Command("uuidgen").Output()

			if err != nil {

				fmt.Println(err)

			}

			m.Uuid = strings.Replace(string(uuid[:]), "\n", "", -1)

			m.Id = 1

			m.Name = "UBER"

			modality := model.Modality{}

			modality.Id = product.ProductID

			if product.DisplayName == "uberBAG" && product.SortDescription == "uberXBAG" {

				modality.Name = product.SortDescription

			} else {

				modality.Name = product.DisplayName

			}

			modality.Image = product.Image

			m.Modality = modality

			m.AlertMessage = player.AlertMessage

			for _, price := range prices.Prices {

				if product.ProductID == price.ProductID {

					m.Price = price.Estimate

					m.Multiplier = price.SurgeMultiplier

					m.PopupMultiplier = "Multiplicador " + strconv.FormatFloat(m.Multiplier, 'f', 1, 64) + "x sobre o valor total da corrida"

					break

				}

			}

			for _, time := range times.Times {

				if time.ProductId == product.ProductID {

					m.WaitingTime = time.Estimate

					break

				}

			}

			response = append(response, m)
		}

	}

	return response
}

func getTimes(start_lat float64, start_lng float64, player model.Player, token model.TokenPlayer) (response model.ResponseTime) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", UBER_DOMAIN+UBER_URL_ESTIMATE_TIME+"?start_latitude="+
		fmt.Sprintf("%f", start_lat)+"&start_longitude="+fmt.Sprintf("%f", start_lng), nil)

	if err != nil {

		return model.ResponseTime{}

	}

	req.Header.Add(UBER_HEADER_AUTH, token.Token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return response
	}

	defer resp.Body.Close()

	htmlData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return response
	}

	b := []byte(string(htmlData))

	var m model.ResponseTime

	err = json.Unmarshal(b, &m)

	return m
}

func getPrices(start_lat float64, start_lng float64, end_lat float64, end_lng float64, player model.Player, token model.TokenPlayer) (response model.ResponsePrices) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", UBER_DOMAIN+UBER_URL_ESTIMATE_PRICE+"?start_latitude="+
		fmt.Sprintf("%f", start_lat)+"&start_longitude="+fmt.Sprintf("%f", start_lng)+
		"&end_latitude="+fmt.Sprintf("%f", end_lat)+"&end_longitude="+fmt.Sprintf("%f", end_lng), nil)

	if err != nil {

		return model.ResponsePrices{}

	}

	req.Header.Add(UBER_HEADER_AUTH, token.Token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return response
	}

	defer resp.Body.Close()

	htmlData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return response
	}

	b := []byte(string(htmlData))

	var m model.ResponsePrices

	err = json.Unmarshal(b, &m)

	return m
}

func getProducts(start_lat float64, start_lng float64, player model.Player, token model.TokenPlayer) (response model.ResponseProduct) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", UBER_DOMAIN+UBER_URL_PRODUCTS+"?latitude="+
		fmt.Sprintf("%f", start_lat)+"&longitude="+fmt.Sprintf("%f", start_lng), nil)

	if err != nil {

		return model.ResponseProduct{}

	}

	req.Header.Add(UBER_HEADER_AUTH, token.Token)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return response
	}

	defer resp.Body.Close()

	htmlData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return response
	}

	b := []byte(string(htmlData))

	var m model.ResponseProduct

	err = json.Unmarshal(b, &m)

	return m
}
