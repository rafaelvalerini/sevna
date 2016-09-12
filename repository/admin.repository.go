package repository

import( 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
	"agregador/model"
)

func CountEstimates() (cont int64){

	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
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

	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
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

	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
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



