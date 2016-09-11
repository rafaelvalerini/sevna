package service

import (
	"net/http"
	"agregador/model"
	"fmt"
 	"io/ioutil"
 	"os"
 	"encoding/json"
 	"os/exec"
 	"runtime"
    "sync"
    "strings"
)

const (
	UBER_DOMAIN = "https://api.uber.com"
	UBER_URL_PRODUCTS = "/v1/products"
	UBER_URL_ESTIMATE_TIME = "/v1/estimates/time"
	UBER_URL_ESTIMATE_PRICE = "/v1/estimates/price"
	UBER_SERVER_TOKEN = "Token N7fFFJoeenUt06hYaIJ73plRNNZuaXawFTZZ0yVr"
	UBER_HEADER_AUTH = "Authorization"
)

func GetEstimatesUber(start_lat float64, start_lng float64, end_lat float64, end_lng float64) (response []model.Player){
	
	runtime.GOMAXPROCS(3)

	var wg sync.WaitGroup

    wg.Add(3)

    var products model.ResponseProduct

    var times model.ResponseTime

    var prices model.ResponsePrices


    go func() {
        defer wg.Done()

        products = getProducts(start_lat, start_lng)

    }()

	go func() {
        defer wg.Done()

        times = getTimes(start_lat, start_lng)
        
    }()

	go func() {
        defer wg.Done()

        prices = getPrices(start_lat, start_lng, end_lat, end_lng)
        
    }()

    wg.Wait()

	return processEstimates(products, times, prices)

}

func processEstimates(products model.ResponseProduct, times model.ResponseTime, prices model.ResponsePrices) (response []model.Player){
	
	if len(products.Products) > 0 {

		for  _,product := range products.Products {

			m := model.Player{}

			uuid, err := exec.Command("uuidgen").Output()

			if err != nil{

				fmt.Println(err)

		 		os.Exit(1)
		 		
			}

			m.Uuid = strings.Replace(string(uuid[:]),"\n","",-1)

			m.Id = 1;

			m.Name = "UBER"

			modality := model.Modality{}

			modality.Id = product.ProductID

			modality.Name = product.DisplayName

			modality.Image = product.Image

			m.Modality = modality

			for  _,price := range prices.Prices {
				
				if product.ProductID == price.ProductID {

					m.Price = price.Estimate

					m.Multiplier = price.SurgeMultiplier

					break

				}

			}

			for _,time := range times.Times {

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

func getTimes(start_lat float64, start_lng float64) (response model.ResponseTime){
	client := &http.Client{
	}

	req, err := http.NewRequest("GET", UBER_DOMAIN + UBER_URL_ESTIMATE_TIME + "?start_latitude=" + 
		fmt.Sprintf("%f", start_lat) + "&start_longitude="+ fmt.Sprintf("%f", start_lng), nil)

	if err != nil{

		return model.ResponseTime{} 

	}

	req.Header.Add(UBER_HEADER_AUTH, UBER_SERVER_TOKEN)

	resp, err := client.Do(req)

	defer resp.Body.Close()

 	htmlData, err := ioutil.ReadAll(resp.Body)

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	b := []byte(string(htmlData))

	var m model.ResponseTime

	err = json.Unmarshal(b, &m)

	return m
}

func getPrices(start_lat float64, start_lng float64, end_lat float64, end_lng float64) (response model.ResponsePrices){
	client := &http.Client{
	}

	req, err := http.NewRequest("GET", UBER_DOMAIN + UBER_URL_ESTIMATE_PRICE + "?start_latitude=" + 
		fmt.Sprintf("%f", start_lat) + "&start_longitude="+ fmt.Sprintf("%f", start_lng) + 
		"&end_latitude="+ fmt.Sprintf("%f", end_lat) + "&end_longitude="+ fmt.Sprintf("%f", end_lng), nil)

	if err != nil{

		return model.ResponsePrices{} 

	}

	req.Header.Add(UBER_HEADER_AUTH, UBER_SERVER_TOKEN)

	resp, err := client.Do(req)

	defer resp.Body.Close()

 	htmlData, err := ioutil.ReadAll(resp.Body) //<--- here!

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	b := []byte(string(htmlData))

	var m model.ResponsePrices

	err = json.Unmarshal(b, &m)

	return m
}

func getProducts(start_lat float64, start_lng float64) (response model.ResponseProduct){
	client := &http.Client{
	}

	req, err := http.NewRequest("GET", UBER_DOMAIN + UBER_URL_PRODUCTS + "?latitude=" + 
		fmt.Sprintf("%f", start_lat) + "&longitude="+ fmt.Sprintf("%f", start_lng), nil)

	if err != nil{

		return model.ResponseProduct{} 

	}

	req.Header.Add(UBER_HEADER_AUTH, UBER_SERVER_TOKEN)

	resp, err := client.Do(req)

	defer resp.Body.Close()

 	htmlData, err := ioutil.ReadAll(resp.Body) //<--- here!

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	b := []byte(string(htmlData))

	var m model.ResponseProduct

	err = json.Unmarshal(b, &m)

	return m
}