package service

import(
	"agregador/model"
	"strings"
)

func GetModalityByName(modalities []model.Modality, name string) (modality model.Modality){

	for _,mo := range modalities {

		if strings.ToUpper(strings.TrimSpace(mo.Name)) == strings.ToUpper(strings.TrimSpace(name)){

			return mo

		}
		
	}

	return modality

}

func GetPlayer(players []model.Player, id int) (player model.Player){

	for _,mo := range players {

		if mo.Id == id{

			return mo

		}
		
	}

	return player

}