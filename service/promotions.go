package service

import (
	"agregador/model"
	"agregador/repository"
	"fmt"
)

func AddPromotions(players []model.Player) {

 	if(len(players) > 0){

 		promotions := repository.FindAllPromotions()

 		if(len(promotions) > 0){

 			fmt.Println("Tamanho promotion: ",len(promotions))

 			playerfor: for idx,_ := range players {

 				for _,promotion := range promotions {

 					fmt.Println("promotion: ",promotion.Modality,", Modality: ", players[idx].Modality.Name)

 					if(promotion.Modality == players[idx].Modality.Name){

 						fmt.Println(promotion.Modality)

 						players[idx].Modality.Promotion = promotion

 						continue playerfor

 					}

 				}

 			}

 		}

 	}

}