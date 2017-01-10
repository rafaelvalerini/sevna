package service

import (
	"agregador/model"
	"agregador/repository"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func AgregateAllV2(request model.RequestAggregator) (agregator model.Aggregator) {

	agregator = AgregateAll(request)

	var players []model.Player

	var uberX string

	for _, element := range agregator.Players {

		if element.Modality.Name == "uberX" {

			uberX = element.Modality.Id

		}

	}

	for _, element := range agregator.Players {

		if element.Modality.Name == "uberPOOL" {

			element.Modality.Id = uberX

		}

		if element.Id == 1 {

			element.Url = "uber://?product_id=" + element.Modality.Id + "&client_id=eXSLlrv8apu9D9YYsJ-Ly1zKQs99cnGc&action=setPickup&pickup[latitude]=" + agregator.Start.Lat + "&pickup[longitude]=" + agregator.Start.Lng + "&pickup[formatted_address]=" + agregator.Start.Address + "&dropoff[latitude]=" + agregator.End.Lat + "&dropoff[longitude]=" + agregator.End.Lng + "&dropoff[formatted_address]=" + agregator.End.Address

		}

		if element.Id == 2 {

			element.Url = "cabify:///journey?vehicle_type_id=" + element.Modality.Id +
				"&stops[0][name]=" + agregator.Start.Address +
				"&stops[0][loc][latitude]=" + agregator.Start.Lat +
				"&stops[0][loc][longitude]=" + agregator.Start.Lng +
				"&stops[1][name]=" + agregator.End.Address +
				"&stops[1][loc][latitude]=" + agregator.End.Lat +
				"&stops[1][loc][longitude]=" + agregator.End.Lng

		}

		if element.Id == 3 {

			element.Url = "taxis99://call?startLat=" + agregator.Start.Lat + "&startLng=" + agregator.Start.Lng +
				"&startName=" + agregator.Start.Address + "&endLat=" + agregator.End.Lat +
				"&endLng=" + agregator.End.Lng + "&endName=" + agregator.End.Address

		}

		if element.Id == 4 {

			element.Url = "easytaxi://p/home"

		}

		elementNotPromo := *element

		elementNotPromo.Modality.Promotion = nil

		elementNotPromo.Modality.Promotions = nil

		players = append(players, elementNotPromo)

		for _, promo := range element.Modality.Promotions {

			elementWithPromo := *element

			element.Modaliy.Promotion = promo

			elementWithPromo = validateAvailableAndCoverage(elementWithPromo)

			if elementWithPromo != nil {

				elementWithPromo.Modality.Promotion = promo

				elementWithPromo.Modality.Promotions = nil

				players = append(players, elementWithPromo)

			}

		}
	}

	agregator.Players = players

	return agregator

}

func validateAvailableAndCoverage(element model.Player, origin model.Position) (result model.Player) {

	availableOk := false

	coverageOk := false

	n := time.Now().Weekday()

	if len(element.Modality.Promotion.PromotionAvailable) > 0 {

		for _, a := range element.Modality.Promotion.PromotionAvailable {

			if a.Monday == 1 && n == 1 {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					return
				}

			} else if a.Tuesday == 1 && n == 2 {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					return
				}

			} else if a.Wednesday == 1 && n == 3 {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					return
				}

			} else if a.Thursday == 1 && n == 4 {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					return
				}

			} else if a.Friday == 1 && n == 5 {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					return
				}

			} else if a.Saturday == 1 && n == 6 {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					return
				}

			} else if a.Sunday == 1 && n == 0 {

				isOk := checkHour(a)

				if isOk {
					availableOk = true
					return
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

				if c.City && c.State {

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

		hourActual, _ = strconv.ParseInt(checkHourCompleteZero(d.Hour()) + "" + checkHourCompleteZero(d.Minute()))

		initialHour, _ = strconv.ParseInt(strings.Replace(available.StartHour, ":", "", -1))

		finalHour, _ = strconv.ParseInt(strings.Replace(available.EndHour, ":", "", -1))

		if hourActual >= initialHour && hourActual <= finalHour {

			return true

		} else {

			return false

		}

	}

}

func checkHourCompleteZero(time int) (result string) {

	if time < 10 {

		result = "0" + time

	} else {

		result = "" + time

	}

	return result

}

func AgregateAll(request model.RequestAggregator) (agregator model.Aggregator) {

	uuid, err := exec.Command("uuidgen").Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aggregate := model.Aggregator{
		Start: request.Start,
		End:   request.End,
		Id:    strings.Replace(string(uuid[:]), "\n", "", -1),
	}

	players := repository.FindAllPlayers()

	runtime.GOMAXPROCS(3)

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()

		playerUber := GetPlayer(players, 1)

		ubbers := GetEstimatesUber(request.Start.Lat, request.Start.Lng, request.End.Lat, request.End.Lng, playerUber)

		for _, element := range ubbers {
			aggregate.Players = append(aggregate.Players, element)
		}

	}()

	go func() {

		defer wg.Done()

		playerCabify := GetPlayer(players, 2)

		cabifys := GetEstimatesCabify(request.Start.Lat, request.Start.Lng, request.End.Lat, request.End.Lng, playerCabify)

		for _, element := range cabifys {
			aggregate.Players = append(aggregate.Players, element)
		}

	}()

	go func() {

		defer wg.Done()

		player99 := GetPlayer(players, 3)

		playerEasy := GetPlayer(players, 4)

		defaults := GetEstimates99TaxiAndEasy(request.Start.Lat, request.Start.Lng, request.End.Lat, request.End.Lng, request.Duration, request.Distance, player99, playerEasy)

		for _, element := range defaults {
			aggregate.Players = append(aggregate.Players, element)
		}

	}()

	wg.Wait()

	AddPromotions(aggregate.Players)

	repository.SaveSearch(aggregate, request)

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
