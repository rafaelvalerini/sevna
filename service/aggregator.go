package service

import(
	"agregador/model"
	"agregador/repository"
	"os/exec"
	"fmt"
	"os"
	"strings"
	"runtime"
	"sync"
	
)

func AgregateAll(request model.RequestAggregator) (agregator model.Aggregator){
	
	uuid, err := exec.Command("uuidgen").Output()

	if err != nil{
		fmt.Println(err)
 		os.Exit(1)
	}

	aggregate := model.Aggregator{
		Start: request.Start,
		End: request.End,
		Id: strings.Replace(string(uuid[:]),"\n","",-1),
	}

	players := repository.FindAllPlayers()

	runtime.GOMAXPROCS(3)

	var wg sync.WaitGroup

    wg.Add(3)

 	go func() {
        defer wg.Done()

        playerUber := GetPlayer(players, 1);

		ubbers := GetEstimatesUber(request.Start.Lat, request.Start.Lng, request.End.Lat, request.End.Lng, playerUber)

		for _,element := range ubbers {
			aggregate.Players = append(aggregate.Players, element)
		}

	}()

	go func() {

        defer wg.Done()

        playerCabify := GetPlayer(players, 2);

		cabifys := GetEstimatesCabify(request.Start.Lat, request.Start.Lng, request.End.Lat, request.End.Lng, playerCabify)

		for _,element := range cabifys {
			aggregate.Players = append(aggregate.Players, element)
		}

	}()

	go func() {

        defer wg.Done()

        player99 := GetPlayer(players, 3)

        playerEasy := GetPlayer(players, 4)

		defaults := GetEstimates99TaxiAndEasy(request.Start.Lat, request.Start.Lng, request.End.Lat, request.End.Lng, request.Duration, request.Distance, player99, playerEasy)

		for _,element := range defaults {
			aggregate.Players = append(aggregate.Players, element)
		}

	}()

	wg.Wait()

	AddPromotions(aggregate.Players)

	go func() {

		repository.SaveSearch(aggregate, request)

	}()

	return aggregate
}


func GetModality(modalities []model.Modality, name string) (modality model.Modality){

	for _,mo := range modalities {

		if mo.Name == name{

			return mo

		}
		
	}

	return modality

}