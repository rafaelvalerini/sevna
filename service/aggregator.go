package service

import (
	"agregador/model"
	"agregador/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

func AgregateAllV1(request model.RequestAggregator) (agregator model.Aggregator) {

	agregator = AgregateAll(request, false)

	repository.SaveSearch(agregator, request)

	return agregator

}

func AgregateAllV21(request model.RequestAggregator) (agregator model.Aggregator) {

	return AgregateAllV2(request, true)

}

func AgregateAllV2(request model.RequestAggregator, yetgo bool) (agregator model.Aggregator) {

	agregator = AgregateAll(request, yetgo)

	var players []model.Player

	var uberX string

	for _, element := range agregator.Players {

		if element.Modality.Name == "uberX" {

			uberX = element.Modality.Id

		}

	}

	for _, element := range agregator.Players {

		var buffer bytes.Buffer

		if element.Modality.Name == "uberPOOL" {

			element.Modality.Id = uberX

		}

		if element.Id == 1 {

			buffer.WriteString("uber://?product_id=")
			buffer.WriteString(element.Modality.Id)
			buffer.WriteString("&client_id=eXSLlrv8apu9D9YYsJ-Ly1zKQs99cnGc&action=setPickup&pickup[latitude]=")
			buffer.WriteString(strconv.FormatFloat(agregator.Start.Lat, 'G', -1, 64))
			buffer.WriteString("&pickup[longitude]=")
			buffer.WriteString(strconv.FormatFloat(agregator.Start.Lng, 'G', -1, 64))
			buffer.WriteString("&pickup[formatted_address]=")
			buffer.WriteString(agregator.Start.Address)
			buffer.WriteString("&dropoff[latitude]=")
			buffer.WriteString(strconv.FormatFloat(agregator.End.Lat, 'G', -1, 64))
			buffer.WriteString("&dropoff[longitude]=")
			buffer.WriteString(strconv.FormatFloat(agregator.End.Lng, 'G', -1, 64))
			buffer.WriteString("&dropoff[formatted_address]=")
			buffer.WriteString(agregator.End.Address)

			element.Url = buffer.String()

			element.UrlAndroid = buffer.String()

		}

		if element.Id == 2 {

			buffer.WriteString("cabify:///journey?vehicle_type_id=")
			buffer.WriteString(element.Modality.Id)
			buffer.WriteString("&stops[0][name]=")
			buffer.WriteString(agregator.Start.Address)
			buffer.WriteString("&stops[0][loc][latitude]=")
			buffer.WriteString(strconv.FormatFloat(agregator.Start.Lat, 'G', -1, 64))
			buffer.WriteString("&stops[0][loc][longitude]=")
			buffer.WriteString(strconv.FormatFloat(agregator.Start.Lng, 'G', -1, 64))
			buffer.WriteString("&stops[1][name]=")
			buffer.WriteString(agregator.End.Address)
			buffer.WriteString("&stops[1][loc][latitude]=")
			buffer.WriteString(strconv.FormatFloat(agregator.End.Lat, 'G', -1, 64))
			buffer.WriteString("&stops[1][loc][longitude]=")
			buffer.WriteString(strconv.FormatFloat(agregator.End.Lng, 'G', -1, 64))

			element.Url = buffer.String()

			element.UrlAndroid = "https://app.adjust.com/a0kvh4?deep_link=" + url.QueryEscape(buffer.String())

		}

		if element.Id == 3 {

			buffer.WriteString("taxis99://call?startLat=")
			buffer.WriteString(strconv.FormatFloat(agregator.Start.Lat, 'G', -1, 64))
			buffer.WriteString("&startLng=")
			buffer.WriteString(strconv.FormatFloat(agregator.Start.Lng, 'G', -1, 64))
			buffer.WriteString("&startName=")
			buffer.WriteString(agregator.Start.Address)
			buffer.WriteString("&endLat=")
			buffer.WriteString(strconv.FormatFloat(agregator.End.Lat, 'G', -1, 64))
			buffer.WriteString("&endLng=")
			buffer.WriteString(strconv.FormatFloat(agregator.End.Lng, 'G', -1, 64))
			buffer.WriteString("&endName=")
			buffer.WriteString(agregator.End.Address)

			element.Url = buffer.String()

			element.UrlAndroid = "https://app.adjust.com/gvl80l_8y8218?deep_link=" + url.QueryEscape(buffer.String())
		}

		if element.Id == 4 {

			element.Url = "easytaxi://p/home"

			element.UrlAndroid = "easytaxi://p/home"

		}

		if element.Id == 5 {

			buffer.WriteString("yetgo://content/?action=setPickup&departurelatitude")
			buffer.WriteString(strconv.FormatFloat(agregator.Start.Lat, 'G', -1, 64))
			buffer.WriteString("&departurelongitude=")
			buffer.WriteString(strconv.FormatFloat(agregator.Start.Lng, 'G', -1, 64))
			buffer.WriteString("&destinationlatitude=")
			buffer.WriteString(strconv.FormatFloat(agregator.End.Lat, 'G', -1, 64))
			buffer.WriteString("&destinationlongitude=")
			buffer.WriteString(strconv.FormatFloat(agregator.End.Lng, 'G', -1, 64))

			element.Url = buffer.String()

			element.UrlAndroid = buffer.String()

		}

		elementNotPromo := createCopyPlayer(element)

		var pro model.Promotion

		elementNotPromo.Modality.Promotion = pro

		elementNotPromo.Modality.Promotions = nil

		arrayPrice := strings.Split(strings.Replace(strings.Replace(elementNotPromo.Price, "R$", "", -1), ",", ".", -1), "-")

		if len(arrayPrice) > 1 {

			priceFirst, _ := strconv.ParseFloat(strings.Replace(arrayPrice[0], ",", ".", -1), 64)

			priceLast, _ := strconv.ParseFloat(strings.Replace(arrayPrice[1], ",", ".", -1), 64)

			elementNotPromo.Index = ((priceFirst + priceLast) / float64(2))

		} else {

			elementNotPromo.Index, _ = strconv.ParseFloat(strings.Replace(arrayPrice[0], ",", ".", -1), 64)

		}

		players = append(players, elementNotPromo)

		for _, promo := range element.Modality.Promotions {

			elementWithPromo := createCopyPlayer(element)

			uuid, _ := exec.Command("uuidgen").Output()

			elementWithPromo.Uuid = strings.Replace(string(uuid[:]), "\n", "", -1)

			elementWithPromo.Modality.Promotion = promo

			elementWithPromo = validateAvailableAndCoverage(elementWithPromo, agregator.Start)

			if elementWithPromo.Id > 0 {

				elementWithPromo.Modality.Promotion = promo

				elementWithPromo.Modality.Promotions = nil

				players = append(players, applyPromotion(elementWithPromo))

			}

		}
	}

	var plyrs model.Players

	plyrs = players

	sort.Sort(plyrs)

	agregator.Players = plyrs

	repository.SaveSearch(agregator, request)

	return agregator

}

func createCopyPlayer(player model.Player) (newPlayer model.Player) {

	b, _ := json.Marshal(player)

	json.Unmarshal(b, &newPlayer)

	return newPlayer

}

func applyPromotion(player model.Player) (playerTmp model.Player) {

	arrayPrice := strings.Split(strings.Replace(player.Price, "R$", "", -1), "-")

	price, _ := strconv.ParseFloat(strings.Replace(arrayPrice[0], ",", ".", -1), 64)

	priceFinal := float64(0)

	if len(arrayPrice) > 1 {

		priceFinal, _ = strconv.ParseFloat(strings.Replace(arrayPrice[1], ",", ".", -1), 64)

		offFinal := (float64(player.Modality.Promotion.Off) / float64(100) * priceFinal)

		if player.Modality.Promotion.LimitOff > 0 {

			if offFinal > player.Modality.Promotion.LimitOff {

				priceFinal = priceFinal - player.Modality.Promotion.LimitOff

			} else {

				priceFinal = priceFinal - offFinal

			}

		}

	}

	var off = (player.Modality.Promotion.Off / 100 * price)

	if player.Modality.Promotion.LimitOff > 0 {

		if off > player.Modality.Promotion.LimitOff {

			price = price - player.Modality.Promotion.LimitOff

		} else {

			price = price - off

		}

	}

	if priceFinal > 0 {

		player.Price = "R$" + strconv.Itoa(int(price)) + "-" + strconv.Itoa(int(priceFinal)+1)

		player.Index = float64((price + (priceFinal + float64(1))) / float64(2))

	} else {

		player.Price = "R$" + strconv.Itoa(int(price)) + "-" + strconv.Itoa(int(price)+1)

		player.Index = float64((price + (price + float64(1))) / float64(2))

	}

	return player

}

func validateAvailableAndCoverage(element model.Player, origin model.Position) (result model.Player) {

	availableOk := false

	coverageOk := false

	n := time.Now().Weekday()

	if len(element.Modality.Promotion.PromotionAvailable) > 0 {

		for _, a := range element.Modality.Promotion.PromotionAvailable {

			if a.Monday == 1 && n.String() == "Monday" {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					break
				}

			} else if a.Tuesday == 1 && n.String() == "Tuesday" {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					break
				}

			} else if a.Wednesday == 1 && n.String() == "Wednesday" {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					break
				}

			} else if a.Thursday == 1 && n.String() == "Thursday" {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					break
				}

			} else if a.Friday == 1 && n.String() == "Friday" {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					break
				}

			} else if a.Saturday == 1 && n.String() == "Saturday" {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					break
				}

			} else if a.Sunday == 1 && n.String() == "Sunday" {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					break
				}

			}
		}

	} else {

		availableOk = true

	}

	if availableOk {

		if len(element.Modality.Promotion.PromotionCoverages) <= 0 {

			coverageOk = true

		} else {

			for _, c := range element.Modality.Promotion.PromotionCoverages {

				if strings.TrimSpace(c.City) != "" && strings.TrimSpace(c.State) != "" {

					if strings.TrimSpace(c.City) == strings.TrimSpace(origin.City) && strings.TrimSpace(c.State) == strings.TrimSpace(origin.State) {

						coverageOk = true

					}

				} else if strings.TrimSpace(c.City) == "" && strings.TrimSpace(c.State) != "" {

					if strings.TrimSpace(c.State) == strings.TrimSpace(origin.State) {

						coverageOk = true

					}

				} else if strings.TrimSpace(c.City) == "" && strings.TrimSpace(c.State) == "" {

					coverageOk = true

				}
			}

		}

	}

	if availableOk && coverageOk {

		return element

	} else {

		return result

	}

}

func checkHour(available model.Available) (isOk bool) {

	if available.StartHour == "" || available.EndHour == "" {

		return true

	} else {

		var d = time.Now()

		hourActual, _ := strconv.Atoi(checkHourCompleteZero(d.Hour()) + "" + checkHourCompleteZero(d.Minute()))

		initialHour, _ := strconv.Atoi(strings.Replace(available.StartHour, ":", "", -1))

		finalHour, _ := strconv.Atoi(strings.Replace(available.EndHour, ":", "", -1))

		if hourActual >= initialHour && hourActual <= finalHour {

			return true

		} else {

			return false

		}

	}

}

func checkHourCompleteZero(time int) (result string) {

	if time < 10 {

		result = "0" + strconv.Itoa(time)

	} else {

		result = "" + strconv.Itoa(time)

	}

	return result

}

func AgregateAll(request model.RequestAggregator, yetgo bool) (agregator model.Aggregator) {

	uuid, err := exec.Command("uuidgen").Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aggregate := model.Aggregator{
		Start:    request.Start,
		End:      request.End,
		Id:       strings.Replace(string(uuid[:]), "\n", "", -1),
		Duration: 0,
		Distance: 0,
	}

	players := repository.FindAllPlayers()

	tokens := repository.GetAccessToken()

	runtime.GOMAXPROCS(4)

	var wg sync.WaitGroup

	wg.Add(4)

	go func() {
		defer wg.Done()

		playerUber := GetPlayer(players, 1)

		tokensUber := GetTokensByPlayer(tokens, 1)

		ubbers := GetEstimatesUber(request.Start.Lat, request.Start.Lng, request.End.Lat, request.End.Lng, playerUber, GetUnicToken(tokensUber))

		for _, element := range ubbers {
			aggregate.Players = append(aggregate.Players, element)
		}

	}()

	go func() {

		defer wg.Done()

		playerCabify := GetPlayer(players, 2)

		tokensCabify := GetTokensByPlayer(tokens, 2)

		cabifys := GetEstimatesCabify(request.Start.Lat, request.Start.Lng, request.End.Lat, request.End.Lng, playerCabify, GetUnicToken(tokensCabify))

		for _, element := range cabifys {
			aggregate.Players = append(aggregate.Players, element)
		}

	}()

	go func() {

		defer wg.Done()

		player99 := GetPlayer(players, 3)

		tokens99 := GetTokensByPlayer(tokens, 3)

		playerEasy := GetPlayer(players, 4)

		defaults := GetEstimates99TaxiAndEasy(request.Start.Lat, request.Start.Lng, request.End.Lat, request.End.Lng, request.Duration, request.Distance, player99, playerEasy, GetUnicToken(tokens99))

		for _, element := range defaults {
			aggregate.Players = append(aggregate.Players, element)
		}

	}()

	if yetgo {
		go func() {
			defer wg.Done()

			playerYetGo := GetPlayer(players, 5)

			tokensYetGo := GetTokensByPlayer(tokens, 5)

			yetts := GetEstimatesYetGo(request.Start.Lat, request.Start.Lng, request.End.Lat, request.End.Lng, playerYetGo, GetUnicToken(tokensYetGo))

			for _, element := range yetts {
				aggregate.Players = append(aggregate.Players, element)
			}

		}()
	} else {
		go func() {
			defer wg.Done()

		}()
	}

	wg.Wait()

	AddPromotions(aggregate.Players)

	return aggregate
}

func GetModality(modalities []model.Modality, name string) (modality model.Modality) {

	for _, mo := range modalities {

		if mo.Name == name {

			return mo

		}

	}

	return modality

}
