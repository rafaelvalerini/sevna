package service

import (
	"agregador/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

const (
	CABIFY_DOMAIN         = "https://cabify.com"
	CABIFY_URL_ESTIMATE   = "/api/v2/estimate"
	CABIFY_HEADER_AUTH    = "Authorization"
	CABIFY_HEADER_CONTENT = "Content-Type"
	CABIFY_CONTENT_JSON   = "application/json"
)

func GetEstimatesCabify(start_lat float64, start_lng float64, end_lat float64, end_lng float64, player model.Player) (response []model.Player) {

	if player.Active == 1 {

		estimate := getEstimates(start_lat, start_lng, end_lat, end_lng, player)

		return processEstimatesCabify(estimate, player)

	} else {

		return response

	}

}

func processEstimatesCabify(estimate model.ResponseCabify, player model.Player) (response []model.Player) {

	time := 420

	for _, est := range estimate {

		modal := GetModalityByName(player.Modalities, est.VehicleType.Name, "")

		if modal.Name == "" || modal.Active == 0 {

			continue

		}

		time = time + 60

		m := model.Player{}

		uuid, _ := exec.Command("uuidgen").Output()

		m.Uuid = strings.Replace(string(uuid[:]), "\n", "", -1)

		m.Id = 2

		m.Name = "CABIFY"

		modality := model.Modality{}

		modality.Id = strings.TrimSpace(est.VehicleType.ID)

		modality.Name = strings.TrimSpace(est.VehicleType.Name)

		modality.Image = strings.TrimSpace(est.VehicleType.Icons.Regular)

		m.Modality = modality

		m.Price = est.FormattedPrice

		m.WaitingTime = time

		m.AlertMessage = player.AlertMessage

		response = append(response, m)

	}

	return response

}

func getEstimates(start_lat float64, start_lng float64, end_lat float64, end_lng float64, player model.Player) (response model.ResponseCabify) {

	client := &http.Client{}

	request := model.RequestCabify{Stop: []model.Stop{model.Stop{Loc: []float64{start_lat, start_lng}}, model.Stop{Loc: []float64{end_lat, end_lng}}}}

	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(request)

	req, err := http.NewRequest("POST", CABIFY_DOMAIN+CABIFY_URL_ESTIMATE, b)

	if err != nil {

		return model.ResponseCabify{}

	}

	req.Header.Add(CABIFY_HEADER_AUTH, player.Token)

	req.Header.Add(CABIFY_HEADER_CONTENT, CABIFY_CONTENT_JSON)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return model.ResponseCabify{}
	}

	htmlData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return response
	}

	c := []byte(string(htmlData))

	var m model.ResponseCabify

	err = json.Unmarshal(c, &m)

	if err != nil {
		fmt.Println(err)
		return response
	}

	return m
}
