package service

import (
	"agregador/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
)

const (
	YETGO_DOMAIN         = "http://app.yetgo.com.br"
	YETGO_URL_ESTIMATE   = "/api/v2/estimate/vah"
	YETGO_HEADER_CONTENT = "Content-Type"
	YETGO_CONTENT_JSON   = "application/json"
)

func GetEstimatesYetGo(start_lat float64, start_lng float64, end_lat float64, end_lng float64, player model.Player, token model.TokenPlayer) (response []model.Player) {

	if player.Active == 1 {

		estimate := getEstimatesYepGo(start_lat, start_lng, end_lat, end_lng, player, token)

		return processEstimatesYetGo(estimate, player)

	} else {

		return response

	}

}

func processEstimatesYetGo(estimate model.ResponseYetGo, player model.Player) (response []model.Player) {

	time := 420

	for _, est := range estimate.Data.Categories {

		modal := GetModalityByName(player.Modalities, est.DisplayName, "")

		if modal.Name == "" || modal.Active == 0 || (est.DisplayName == "COMUM" && est.SubcategoryDisplayName != "COMUM") {

			continue

		}

		time = int(estimate.Data.MinArrivalTime) * 60

		m := model.Player{}

		uuid, _ := exec.Command("uuidgen").Output()

		m.Uuid = strings.Replace(string(uuid[:]), "\n", "", -1)

		m.Id = 5

		m.Name = "YETGO"

		modality := model.Modality{}

		modality.Id = strings.TrimSpace(est.DisplayName)

		modality.Name = strings.TrimSpace(modal.Name)

		m.Modality = modality

		m.Price = "R$" + est.EstimatedPrice

		m.WaitingTime = time

		m.AlertMessage = player.AlertMessage

		response = append(response, m)

	}

	return response

}

func getEstimatesYepGo(start_lat float64, start_lng float64, end_lat float64,
	end_lng float64, player model.Player, token model.TokenPlayer) (response model.ResponseYetGo) {

	client := &http.Client{}

	form := url.Values{}
	form.Add("start_latitude", strconv.FormatFloat(start_lat, 'G', -1, 64))
	form.Add("start_longitude", strconv.FormatFloat(start_lng, 'G', -1, 64))
	form.Add("destination_latitude", strconv.FormatFloat(end_lat, 'G', -1, 64))
	form.Add("destination_longitude", strconv.FormatFloat(end_lng, 'G', -1, 64))
	form.Add("api_key", token.Token)

	req, err := http.NewRequest("POST", YETGO_DOMAIN+YETGO_URL_ESTIMATE, strings.NewReader(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {

		return response

	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return response
	}

	htmlData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return response
	}

	c := []byte(string(htmlData))

	var m model.ResponseYetGo

	err = json.Unmarshal(c, &m)

	if err != nil {
		fmt.Println(err)
		return response
	}

	return m
}
