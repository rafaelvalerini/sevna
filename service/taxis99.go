package service

import (
	"net/http"
	"agregador/model"
	"fmt"
 	"io/ioutil"
 	"encoding/json"
    "os/exec"
    "bytes"
    "strings"
)

const (
	TAXIS99_DOMAIN = "https://api.99taxis.com"
	TAXIS99_URL_ESTIMATE = "/v1/categories/pricingEstimates"
	TAXIS99_HEADER_USER_ID = "X-User-Id"
	TAXIS99_HEADER_USER_ID_VALUE = "10878359"
	TAXIS99_HEADER_USER_UUID = "X-User-UUID"
	TAXIS99_HEADER_USER_UUID_VALUE = ""
	TAXIS99_HEADER_CONTENT = "Content-Type" 
	TAXIS99_CONTENT_JSON = "application/json"
)

func GetEstimates99TaxiAndEasy(start_lat float64, start_lng float64, end_lat float64, end_lng float64, duration int64, distance int64, player99 model.Player, playerEasy model.Player) (response []model.Player){

    estimate := getEstimates99(start_lat, start_lng, end_lat, end_lng, duration, distance, player99)

	return processEstimates99Taxi(estimate, player99, playerEasy)

}

func processEstimates99Taxi(estimate model.Response99Taxi, player99 model.Player, playerEasy model.Player) (response []model.Player){
	
	time := 420

	for  _,est := range estimate.Estimates {

		time = time + 60

		m := model.Player{}

		uuid,_ := exec.Command("uuidgen").Output()

		m.Uuid = strings.Replace(string(uuid[:]),"\n","",-1)

		m.Id = 3;

		m.Name = "99 TAXI"

		modality := model.Modality{}

		uuid2,_ := exec.Command("uuidgen").Output()

		modality.Id = strings.Replace(string(uuid2[:]),"\n","",-1)

		switch est.CategoryID{
		    case "pop99" :
		        modality.Name = "99POP"
		        m.WaitingTime = 420
		        break
		    case "regular-taxi":
		        modality.Name = "Táxi"
		        m.WaitingTime = 420
		        break
		    case "top99":
		        modality.Name = "99TOP"
		        m.WaitingTime = 300
		        break
		   	case "turbo-taxi":
		        modality.Name = "Táxi 30% OFF"
		        m.WaitingTime = 240
		        break
	    }

	    modal := GetModalityByName(player99.Modalities, modality.Name)

		if modal.Name == "" || modal.Active == 0 || player99.Active == 0{

			continue

		}

		m.Modality = modal

		m.Price = "R$" + est.LowerFare + "-" + est.UpperFare

		response = append(response, m)

		playerEasyReturn := getEasy(m, playerEasy)

		if playerEasyReturn.Id == 0{

			continue

		}else{

			response = append(response, playerEasyReturn);

		}

	}

	return response

}

func getEasy(m model.Player, playerEasy model.Player) (playerReturn model.Player){
	
	playerResult := model.Player{}

	playerResult.Id = 4

	playerResult.Name = "EASY TAXI"

	uuid,_ := exec.Command("uuidgen").Output()

	playerResult.Uuid = strings.Replace(string(uuid[:]),"\n","",-1)

	uuid2,_ := exec.Command("uuidgen").Output()

	playerResult.Modality.Id = strings.Replace(string(uuid2[:]),"\n","",-1)

	playerResult.Price = m.Price

	switch m.Modality.Name{
	    case "99 POP" :
	        playerResult.Modality.Name = "Easy Go"
	        playerResult.WaitingTime = 300
	        break
	    case "99 TÁXI":
	        playerResult.Modality.Name = "EasyTaxi"
	        playerResult.WaitingTime = 360
	        break
	    case "99 TOP":
	        playerResult.Modality.Name = "EasyPlus+"
	        playerResult.WaitingTime = 300
	        break
	   	case "99 TÁXI 30% off":
	        playerResult.Modality.Name = "EasyTaxi 30% OFF"
	        playerResult.WaitingTime = 240
	        break
	}

	modal := GetModalityByName(playerEasy.Modalities, playerResult.Modality.Name)

	if modal.Name == "" || modal.Active == 0 || playerEasy.Active == 0{

		return playerReturn

	}

	return playerResult
}

func getEstimates99(start_lat float64, start_lng float64, end_lat float64, end_lng float64, duration int64, distance int64, player99 model.Player) (response model.Response99Taxi){
	
	client := &http.Client{}

	request := model.Request99Taxi{
		DistanceInMeters: distance, 
		TimeInSeconds: (duration * 60), 
		PickupLatitude: start_lat, 
		PickupLongitude: start_lng, 
		DropoffLatitude: end_lat, 
		DropoffLongitude: end_lng,
	}

	b := new(bytes.Buffer)

    json.NewEncoder(b).Encode(request)

	req, err := http.NewRequest("POST", TAXIS99_DOMAIN + TAXIS99_URL_ESTIMATE , b)

	if err != nil{

		return model.Response99Taxi{}

	}

	req.Header.Add(TAXIS99_HEADER_USER_ID, TAXIS99_HEADER_USER_ID_VALUE)

	req.Header.Add(TAXIS99_HEADER_USER_UUID, player99.Token)

	req.Header.Add(TAXIS99_HEADER_CONTENT, CABIFY_CONTENT_JSON)

	resp, err := client.Do(req)

	if err != nil {
 		return model.Response99Taxi{}
 	}

	defer resp.Body.Close()

 	htmlData, err := ioutil.ReadAll(resp.Body) 

 	if err != nil {
 		fmt.Println(err)
 	}

 	c := []byte(string(htmlData))

	var m model.Response99Taxi

	err = json.Unmarshal(c, &m)

	if err != nil {
 		fmt.Println(err)
 	}

	return m
}