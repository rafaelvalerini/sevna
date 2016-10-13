package repository

import( 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
	"agregador/model"
    "time"
    "strconv"
)

func CountEstimates() (cont int64){

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err)

	}

	rows, err := db.Query("select count(id) as cont from search") 

    if err != nil {

       fmt.Println(err)

        os.Exit(1)

    }

    for rows.Next() {

        err = rows.Scan(&cont)

    }
    
    defer db.Close()

    return cont

}

func CountModalities() (modalities []model.MoreUser){

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err)

	}

	rows, err := db.Query("select s.modality as modality, count(*) as cont" +
							" from search_results s  " +
								" inner join search_selected ss on s.id = ss.id_search_results " +
							" group by s.modality") 

    if err != nil {

       fmt.Println(err)

        os.Exit(1)

    }

    for rows.Next() {

    	var modality string

    	var count int64

        err = rows.Scan(&modality, &count)

        element := model.MoreUser{
                Modality: modality, 
                Value: count,
            }

        modalities = append(modalities, element)

    }
    
    defer db.Close()

    return modalities

}

func SearchByNotification(state string, city string, initial int) (notifications []string){

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err)

	}

	rows, err := db.Query(" select distinct device.id_notification " +
							" from config_device device  " +
								" inner join search s on s.id = device.id_search  " +
								" inner join search_address address on (address.id = s.start_address_id or address.id = s.end_address_id)  " +
							" where ( = '' or ? = address.state) and (? = '' or ? = address.city)  " +
							" order by device.id_notification  " +
							" LIMIT ?,100", state, state, city, city, initial) 

    if err != nil {

       fmt.Println(err)

        os.Exit(1)

    }

    for rows.Next() {

    	var device string

        err = rows.Scan(&device)

        notifications = append(notifications, device)

    }
    
    defer db.Close()

    return notifications

}


func SumSavedEstimates() (model model.MetaFloat64){

    db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
    
    if err != nil {

        panic(err)

    }

    rows, err := db.Query("select "+
                            "sum( tot.soma ) "+
                        "from "+
                            "( "+
                                "select "+
                                    "value_max.id, "+
                                    "value_max.count_max, "+
                                    "sel.count_min, "+
                                    "( "+
                                        "value_max.count_max - sel.count_min "+
                                    ") as soma "+
                                "from "+
                                    "( "+
                                        "select "+
                                           "s.id as id, "+
                                            "max( cast( SUBSTRING_INDEX( REPLACE( REPLACE( REPLACE( REPLACE( REPLACE( REPLACE( sr.tax_value, 'N/A', '10' ), 'Unavailable', '10' ), '€', '' ), ' ', '' ), ',', '.' ), 'R$', '' ), '-', 1 ) as decimal( 8, 2 ) ) ) as count_max "+
                                        "from "+
                                            "search s inner join search_results sr on "+
                                            "s.id = sr.id_search "+
                                        "group by "+
                                            "s.id "+
                                    ") value_max left join( "+
                                        "select "+
                                            "sr.id_search as id, "+
                                            "min( cast( SUBSTRING_INDEX( REPLACE( REPLACE( REPLACE( REPLACE( REPLACE( REPLACE( sr.tax_value, 'N/A', '10' ), 'Unavailable', '10' ), '€', '' ), ' ', '' ), ',', '.' ), 'R$', '' ), '-', 1 ) as decimal( 8, 2 ) ) ) as count_min "+
                                        "from "+
                                            "search_results sr inner join search_selected ss on "+
                                            "sr.id = ss.id_search_results "+
                                       " group by "+
                                            "sr.id_search "+
                                    ") sel on "+
                                    "value_max.id = sel.id "+
                            ") tot") 

    if err != nil {

       fmt.Println(err)

        os.Exit(1)

    }

    for rows.Next() {

        err = rows.Scan(&model.Value)

    }
    
    defer db.Close()

    return model

}

func FindAnalytics(state string, city string, player string, modality string, startAt string, endAt string) (result []model.ResultAnalytics){

    db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
    
    if err != nil {

        panic(err)

    }

    layout := "2006-01-02T15:04:05.000Z"

    start_date, err := time.Parse(layout, startAt)

    end_date, err := time.Parse(layout, endAt)

    query := fmt.Sprintf("select "+
                "distinct  "+
                "startAddress.date_time as DateTime, "+
                "startAddress.address as StartAddress, "+
                "endAddress.address as EndAddress, "+
                "p.name as Player, "+
                "selected.id_modality as Modality, "+
                "selected.tax_value as Value, "+
                "c.operation_system as OperationSystem, "+
                "startAddress.city as StartCity, "+
                "startAddress.state as StartState, "+
                "endAddress.city as EndCity, "+
                "endAddress.state as EndState, "+
                "c.operation_system_version as OperationSystemVersion, "+
                "c.type_connection as TypeConnection "+
            "from "+
                "( "+
                    "select "+
                        "s.date_time as date_time, "+
                        "s.id as id, "+
                        "sa.address as address, "+
                        "sa.city as city, "+
                        "sa.state as state "+
                    "from "+
                        "search s inner join search_address sa on "+
                        "s.start_address_id = sa.id "+
                ") startAddress left join( "+
                    "select "+
                        "res.id_search as id_search, "+
                        "res.id_player as id_player, "+
                        "res.modality as id_modality, "+
                        "res.tax_value as tax_value "+
                    "from "+
                        "search_results res inner join search_selected se on "+
                        "se.id_search_results = res.id "+
                ") selected on "+
                "selected.id_search = startAddress.id inner join( "+
                    "select "+
                        "s.date_time as date_time, "+
                        "s.id as id, "+
                        "sa.address as address, "+
                        "sa.city as city, "+
                        "sa.state as state "+
                    "from "+
                        "search s inner join search_address sa on "+
                        "s.end_address_id = sa.id "+
                ") endAddress on "+
                "startAddress.id = endAddress.id  "+
                "left join player p on p.id = selected.id_player "+
                "inner join config_device c on startAddress.id = c.id_search "+
            "where startAddress.date_time between '%s' and '%s' ", start_date.Format(time.RFC3339), end_date.Format(time.RFC3339))

    

    if state != ""{
        query = query + " and startAddress.state = '" + state + "'"
    }

    if city != ""{
        query = query + " and startAddress.city = '" + city + "'"
    }

    if player != ""{
        playerId,_ := strconv.Atoi(player)
        query = query + " and p.id = " + fmt.Sprintf("%d",playerId)
    }


    if modality != ""{
        query = query + " and UPPER(selected.id_modality) = '" + modality + "'"
    }

    //query = query + " order by startAddress.date_time desc "

    rows, err := db.Query(query) 

    if err != nil {

       fmt.Println(err)

    }

    for rows.Next() {

        var res model.ResultAnalytics

        var playerResult []byte

        var modalityResult []byte

        var valueResult []byte

        err = rows.Scan(&res.DateTime, &res.StartAddress, &res.EndAddress, &playerResult, &modalityResult, &valueResult, &res.OperationSystem, &res.StartCity, &res.StartState, &res.EndCity, &res.EndState, &res.OperationSystemVersion, &res.TypeConnection)

        res.Player = string(playerResult)

        res.Modality = string(modalityResult)

        res.Value = string(valueResult)

        result = append(result, res)

    }
    
    defer db.Close()

    return result

}