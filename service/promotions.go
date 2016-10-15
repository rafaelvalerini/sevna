package service

import (
	"agregador/model"
	"agregador/repository"
	"strings"
	"strconv"
	"fmt"
)

func AddPromotions(players []model.Player) {

 	if(len(players) > 0){

 		promotions := repository.FindAllPromotions()

 		if(len(promotions) > 0){

 			playerfor: for idx,_ := range players {

 				for _,promotion := range promotions {

 					if promotion.Modality == players[idx].Modality.Name{

 						fmt.Println(players[idx].Modality.Name)

 						if promotion.Id == "3"{

 							value,_ := strconv.ParseFloat(strings.Replace(strings.Replace(strings.Replace(strings.Replace(players[idx].Price, ",", ".", -1), " ", "", -1), "R$", "", -1), "-", "", -1), 64)

 							if value <= 15.0{

 								players[idx].Modality.Promotion = promotion

 								continue playerfor

 							}

 						}else{

 							players[idx].Modality.Promotion = promotion

 							continue playerfor

 						}

 					}

 				}

 			}

 		}

 	}

}