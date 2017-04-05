package service

import (
	"agregador/model"
	"math/rand"
	"strings"
)

func GetModalityByName(modalities []model.Modality, name string, nameApi string) (modality model.Modality) {

	for _, mo := range modalities {

		if strings.ToUpper(strings.TrimSpace(mo.KeyApi)) == strings.ToUpper(strings.TrimSpace(name)) {

			if nameApi != "" {
				mo.NameApi = nameApi
			}

			return mo

		}

	}

	return modality

}

func GetPlayer(players []model.Player, id int) (player model.Player) {

	for _, mo := range players {

		if mo.Id == id {

			return mo

		}

	}

	return player

}

func GetTokensByPlayer(tokens []model.TokenPlayer, playerID int) (tokenByPlayer []model.TokenPlayer) {

	for i := 0; i < len(tokens); i++ {

		if tokens[i].PlayerId == playerID {

			tokenByPlayer = append(tokenByPlayer, tokens[i])

		}

	}

	return tokenByPlayer

}

func GetUnicToken(tokens []model.TokenPlayer) model.TokenPlayer {

	index := rand.Intn(len(tokens))

	return tokens[index]

}
