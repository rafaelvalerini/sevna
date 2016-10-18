package service

import (
	"agregador/model"
	"agregador/repository"
)

func AddPromotions(players []model.Player) {

 	if(len(players) > 0){

 		promotions := repository.FindAllPromotions()

 		if(len(promotions) > 0){

 			playerfor: for idx,_ := range players {

 				for _,promotion := range promotions {

 					if promotion.Modality == players[idx].Modality.Name{

						players[idx].Modality.Promotion = promotion

						continue playerfor

 					}

 				}

 			}

 		}

 	}

}