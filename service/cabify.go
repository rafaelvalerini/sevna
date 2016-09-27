package service

import (
	"net/http"
	"agregador/model"
	"fmt"
 	"io/ioutil"
 	"os"
 	"encoding/json"
    "os/exec"
    "bytes"
    "strings"
)

const (
	//CABIFY_DOMAIN_HML = "https://test.cabify.com"
	//CABIFY_SERVER_TOKEN_HML = "Bearer fLqDw9wj_upqf1FOXSXSJ_UrFL643eSHZwai0M4sBVE"
	CABIFY_DOMAIN = "https://cabify.com"
	CABIFY_URL_ESTIMATE = "/api/v2/estimate"
	CABIFY_SERVER_TOKEN = "Bearer EYa4WcYJ74sN43vTJZUSpYmeNBm7HYVo8hSZQplYvg8"
	CABIFY_HEADER_AUTH = "Authorization"
	CABIFY_HEADER_CONTENT = "Content-Type" 
	CABIFY_CONTENT_JSON = "application/json"
)

func GetEstimatesCabify(start_lat float64, start_lng float64, end_lat float64, end_lng float64) (response []model.Player){

    estimate := getEstimates(start_lat, start_lng, end_lat, end_lng)

	return processEstimatesCabify(estimate)

}

func processEstimatesCabify(estimate model.ResponseCabify) (response []model.Player){
	
	time := 420

	for  _,est := range estimate {

		time = time + 60

		m := model.Player{}

		uuid, _ := exec.Command("uuidgen").Output()

		m.Uuid = strings.Replace(string(uuid[:]),"\n","",-1)

		m.Id = 2;

		m.Name = "CABIFY"

		modality := model.Modality{}

		modality.Id = est.VehicleType.ID

		modality.Name = est.VehicleType.Name

		modality.Image = est.VehicleType.Icons.Regular

		m.Modality = modality

		m.Price = est.FormattedPrice

		m.WaitingTime = time

		response = append(response, m)

	}


	return response
}



func getEstimates(start_lat float64, start_lng float64, end_lat float64, end_lng float64) (response model.ResponseCabify){
	
	client := &http.Client{}

	request := model.RequestCabify{Stop: []model.Stop{model.Stop{Loc: []float64{start_lat, start_lng}}, model.Stop{Loc: []float64{end_lat, end_lng}}}}

	b := new(bytes.Buffer)

    json.NewEncoder(b).Encode(request)

	req, err := http.NewRequest("POST", CABIFY_DOMAIN + CABIFY_URL_ESTIMATE , b)

	if err != nil{

		return model.ResponseCabify{}

	}

	req.Header.Add(CABIFY_HEADER_AUTH, CABIFY_SERVER_TOKEN)
	
	req.Header.Add(CABIFY_HEADER_CONTENT, CABIFY_CONTENT_JSON)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
 		return model.ResponseCabify{}
 	}

 	htmlData, err := ioutil.ReadAll(resp.Body) 

 	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
 	}

 	c := []byte(string(htmlData))

	var m model.ResponseCabify

	err = json.Unmarshal(c, &m)

	return m
}