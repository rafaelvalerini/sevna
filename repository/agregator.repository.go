package repository

import( 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"agregador/model"
	"os/exec"
	"fmt"
	"os"
	"strings"
	"time"
)

func SaveSearch(agregator model.Aggregator, request model.RequestAggregator) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err.Error())

	}

	defer db.Close()

	startAddressId := saveStartAddress(request.Start, db)

	endAddressId := saveStartAddress(request.End, db)

	searchBase := saveSearch(startAddressId, endAddressId, db)	

	if len(agregator.Players) > 0 {

		for _,player := range agregator.Players {
			
			saveResults(searchBase, player, db)

		}

	}

	saveConfigDevice(searchBase, request.Device, db);

}

func saveConfigDevice(searchBase string, device model.Device, db *sql.DB) {
	
	 stmtIns, err := db.Prepare("INSERT INTO config_device (id_device, operation_system, operation_system_version, device, type_connection, id_search) VALUES(?, ?, ?, ?, ?, ?)") 

    if err != nil {

        panic(err.Error())

    }

    _, err = stmtIns.Exec(device.IdDevice, device.OperationSystem, device.OperationSystemVersion, device.Device, device.TypeConnection, searchBase)

    if err != nil {

    	panic(err.Error())

    }
    
    defer stmtIns.Close()

}

func saveResults(searchBase string, player model.Player, db *sql.DB) {
	
    stmtIns, err := db.Prepare("INSERT INTO search_results (id, id_player, modality, waiting_time, tax_value, id_search, multiplier, promotion) VALUES(?, ?, ?, ?, ?, ?, ?, ?)") 

    if err != nil {

        panic(err.Error())

    }

    _, err = stmtIns.Exec(strings.Replace(player.Uuid,"\n","",-1), player.Id, player.Modality.Name, player.WaitingTime, player.Price, searchBase, player.Multiplier, "")

    if err != nil {

    	panic(err.Error())

    }
    
    defer stmtIns.Close()

}

func saveSearch(startAddressId string, endAddressId string, db *sql.DB) (uuidString string){
	
	uuid, err := exec.Command("uuidgen").Output()

	if err != nil{

		fmt.Println(err)

 		os.Exit(1)

	}

    stmtIns, err := db.Prepare("INSERT INTO search (id, date_time, start_address_id, end_address_id) VALUES(?, ?, ?, ?)") 

    if err != nil {

        panic(err.Error())

    }

    uuidString = strings.Replace(string(uuid[:]),"\n","",-1)

    _, err = stmtIns.Exec(uuidString, time.Now(), startAddressId, endAddressId)
    
    if err != nil {

    	panic(err.Error())

    }
    
    defer stmtIns.Close()

    return uuidString

}


func saveStartAddress(position model.Position, db *sql.DB) (uuidString string){
	
	uuid, err := exec.Command("uuidgen").Output()

	if err != nil{

		fmt.Println(err)

 		os.Exit(1)

	}

    stmtIns, err := db.Prepare("INSERT INTO search_address (id, lat, lng, address, district, city, state) VALUES(?, ?, ?, ?, ?, ?, ?)") 

    if err != nil {

        panic(err.Error())

    }

    uuidString = strings.Replace(string(uuid[:]),"\n","",-1)

    _, err = stmtIns.Exec(uuidString, position.Lat, position.Lng, position.Address, position.District, position.City, position.State)
    
    if err != nil {

    	panic(err.Error())

    }
    
    defer stmtIns.Close()

    return uuidString

}


func Selected(selected string) (uuidString string){
	
	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err.Error())

	}

	defer db.Close()

    stmtIns, err := db.Prepare("INSERT INTO search_selected(id_search_results, date_time_click) VALUES(?, ?)") 

    if err != nil {

        panic(err.Error())

    }

    _, err = stmtIns.Exec(selected, time.Now())
    
    if err != nil {

    	panic(err.Error())

    }
    
    defer stmtIns.Close()

    return selected

}