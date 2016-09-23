package service

import (
	"agregador/model"
	"agregador/repository"
    "os/exec"
    "strings"
    "fmt"
)

func GetEstimatesDefault(request model.RequestAggregator) (response []model.Player){

    estimate := getEstimatesDefault(request.Start)

	return processEstimatesDefault(request.Distance, estimate)

}

func processEstimatesDefault(distance int64, estimate []model.Player) (response []model.Player){
	
	for  _,est := range estimate {

		uuid, _ := exec.Command("uuidgen").Output()

		est.Uuid = strings.Replace(string(uuid[:]),"\n","",-1)

		price := float64(distance/1000) * est.Modality.PriceKm

		est.Price = fmt.Sprintf("R$%.0f-%.0f", price, price + (price * 20 / 100))

		est.WaitingTime = est.Modality.TimeKm

		est.Multiplier = 1

		response = append(response, est)

	}

	return response
}



func getEstimatesDefault(position model.Position) (response []model.Player){
	
	return repository.SearchPlayersDefault(position);

}