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

	runtime.GOMAXPROCS(3)

	var wg sync.WaitGroup

    wg.Add(3)


 	go func() {
        defer wg.Done()

		ubbers := GetEstimatesUber(request.Start.Lat, request.Start.Lng, request.End.Lat, request.End.Lng)

		for _,element := range ubbers {
			aggregate.Players = append(aggregate.Players, element)
		}

	}()

	go func() {

        defer wg.Done()

		cabifys := GetEstimatesCabify(request.Start.Lat, request.Start.Lng, request.End.Lat, request.End.Lng)

		for _,element := range cabifys {
			aggregate.Players = append(aggregate.Players, element)
		}

	}()

	go func() {

        defer wg.Done()

		defaults := GetEstimatesDefault(request)

		for _,element := range defaults {
			aggregate.Players = append(aggregate.Players, element)
		}

	}()

	wg.Wait()

	go func() {

		repository.SaveSearch(aggregate, request)

	}()

	return aggregate
}