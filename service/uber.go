package service

import (
	"agregador/model"
	"agregador/repository"
	"bytes"
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
	UBER_URL_PRODUCTS       = "/v1.2/products"
	UBER_URL_ESTIMATE_PRICE = "/v1.2/requests/estimate"
	UBER_HEADER_AUTH        = "Authorization"
)

func GetEstimatesUber(start_lat float64, start_lng float64, end_lat float64, end_lng float64, player model.Player,
	player99 model.Player, playerEasy model.Player, token model.TokenPlayer, token99 model.TokenPlayer) (response []model.Player, distance int, duration int) {

	if player.Active == 1 {

		tokenUber := repository.FindTokenUber()

		products := getProducts(start_lat, start_lng, player, tokenUber)

		if len(products.Products) > 0 {

			runtime.GOMAXPROCS(len(products.Products) - 1)

			var wg sync.WaitGroup

			wg.Add(len(products.Products) - 1)

			for index, product := range products.Products {

				b, err := json.Marshal(product)
				if err != nil {
					fmt.Println(err)

				} else {

					byt := []byte(string(b))

					var p model.ProductUber

					if err := json.Unmarshal(byt, &p); err != nil {
						panic(err)
					}

					if index == 0 {

						processUberThead(index, p, player, player99, playerEasy, &response, start_lat, start_lng, end_lat, end_lng, tokenUber, token99)

					} else {

						go func() {

							defer wg.Done()

							processUberThead(index, p, player, player99, playerEasy, &response, start_lat, start_lng, end_lat, end_lng, tokenUber, token99)

						}()

					}

				}

			}

			wg.Wait()

		} else {

			response = append(response, getEstimates99AndEasyGoogleMatrix(start_lat, start_lng, end_lat, end_lng, player99, playerEasy, token99)...)

		}

	}

	return response, distance, duration

}

func processUberThead(index int, product model.ProductUber, player model.Player, player99 model.Player, playerEasy model.Player, response *[]model.Player,
	start_lat float64, start_lng float64, end_lat float64, end_lng float64, tokenUber model.TokenUber, token99 model.TokenPlayer) {

	modal := GetModalityByName(player.Modalities, product.DisplayName, "")

	if modal.Name != "" && modal.Active != 0 {

		estimate := getEstimatesUber(start_lat, start_lng, end_lat, end_lng, tokenUber, product.ProductID)

		if index == 0 {

			if estimate.Trip.DistanceEstimate > 0 {

				distance := int(estimate.Trip.DistanceEstimate * 1000 * 1.6)

				duration := int(estimate.Trip.DurationEstimate / 60)

				players99AndEasy := GetEstimates99TaxiAndEasy(start_lat, start_lng, end_lat, end_lng, int64(duration), int64(distance), player99, playerEasy, token99)

				*response = append(*response, players99AndEasy...)

			} else {

				players99AndEasy := getEstimates99AndEasyGoogleMatrix(start_lat, start_lng, end_lat, end_lng, player99, playerEasy, token99)

				*response = append(*response, players99AndEasy...)

			}

		}

		if estimate.Trip.DistanceEstimate <= 0 {
			return
		}

		playerAdd := processEstimates(player, product.ProductID, product.ProductGroup, product.DisplayName, product.ShortDescription, estimate, product.Image)

		if playerAdd.Id > 0 {

			*response = append(*response, playerAdd)

		}

	}

}

func getEstimates99AndEasyGoogleMatrix(start_lat float64, start_lng float64, end_lat float64, end_lng float64,
	player99 model.Player, playerEasy model.Player,
	token model.TokenPlayer) (response []model.Player) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://maps.googleapis.com/maps/api/distancematrix/json?key=AIzaSyDa2yKVjlQEGzrtwdwC9Je7evqNyAsiq6s&origins="+strconv.FormatFloat(start_lat, 'f', -1, 64)+
		","+strconv.FormatFloat(start_lng, 'f', -1, 64)+"&destinations="+strconv.FormatFloat(end_lat, 'f', -1, 64)+
		","+strconv.FormatFloat(end_lng, 'f', -1, 64), nil)

	if err != nil {

		fmt.Println(err)

		return response

	}

	req.Header.Add("Content-Type", "application/json")

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

	c := []byte(string(htmlData))

	var m model.ResponseGoogleMatrix

	err = json.Unmarshal(c, &m)

	if err != nil {
		fmt.Println(err)
		return response
	}

	var distante int64

	var duration int64

	if len(m.Rows) > 0 && len(m.Rows[0].Elements) > 0 {

		distante = m.Rows[0].Elements[0].Distance.Value

		duration = m.Rows[0].Elements[0].Duration.Value / 60

	} else {

		return response

	}

	return GetEstimates99TaxiAndEasy(start_lat, start_lng, end_lat, end_lng, duration, distante, player99, playerEasy, token)
}

func processEstimates(player model.Player, productID string, productName string, displayName string,
	shortDescription string, estimate model.ResponseEstimateV12, image string) (response model.Player) {

	modal := GetModalityByName(player.Modalities, displayName, "")

	if modal.Name == "" || modal.Active == 0 {

		return response

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

	modality.Id = productID

	if displayName == "uberBAG" && shortDescription == "uberXBAG" {

		modality.Name = shortDescription

	} else {

		modality.Name = displayName

	}

	modality.Image = image

	m.Modality = modality

	m.AlertMessage = player.AlertMessage

	m.Price = strings.Replace(estimate.Fare.Display, ".", ",", -1)

	m.WaitingTime = estimate.PickupEstimate * 60

	m.Multiplier = float64(1)

	m.PopupMultiplier = "Multiplicador " + strconv.FormatFloat(m.Multiplier, 'f', 1, 64) + "x sobre o valor total da corrida"

	return m
}

func getEstimatesUber(start_lat float64, start_lng float64, end_lat float64, end_lng float64, tokenUber model.TokenUber, productID string) (response model.ResponseEstimateV12) {
	client := &http.Client{}

	request := model.RequestEstimateV12{ProductID: productID, EndLatitude: end_lat, EndLongitude: end_lng,
		SeatCount: strconv.Itoa(2), StartLatitude: start_lat, StartLongitude: start_lng}

	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(request)

	req, err := http.NewRequest("POST", UBER_DOMAIN+UBER_URL_ESTIMATE_PRICE, b)

	if err != nil {

		fmt.Println(err)

		return model.ResponseEstimateV12{}

	}

	req.Header.Add(UBER_HEADER_AUTH, tokenUber.TokenBearer)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Language", "pt_BR")

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

	c := []byte(string(htmlData))

	var m model.ResponseEstimateV12

	err = json.Unmarshal(c, &m)

	return m
}

func getProducts(start_lat float64, start_lng float64, player model.Player, token model.TokenUber) (response model.ResponseProductV12) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", UBER_DOMAIN+UBER_URL_PRODUCTS+"?latitude="+
		fmt.Sprintf("%f", start_lat)+"&longitude="+fmt.Sprintf("%f", start_lng), nil)

	if err != nil {

		fmt.Println(err)

		return model.ResponseProductV12{}

	}

	req.Header.Add(UBER_HEADER_AUTH, token.TokenBearer)

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

	var m model.ResponseProductV12

	err = json.Unmarshal(b, &m)

	return m
}
