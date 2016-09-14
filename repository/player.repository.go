package repository

import( 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
	"agregador/model"
	"strconv"
)

func SavePlayer(entity model.Player) (model model.Player){

	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err.Error())

	}

	defer db.Close()

	id := savePlayer(entity, db)

	saveModality(entity.Modalities, id, db)

	return entity

}

func savePlayer(entity model.Player, db *sql.DB) (id int){
	
	if entity.Id > 0{

		stmtIns, err := db.Prepare("UPDATE player SET name = ? where id = ?");

	    if err != nil {

	        panic(err.Error())

	    }

	    _, err = stmtIns.Exec(entity.Name, entity.Id)

	    if err != nil {

	    	panic(err.Error())

	    }
	       
	    defer stmtIns.Close()

	    return entity.Id

	}else{

		stmtIns, err := db.Prepare("INSERT INTO player (name) VALUES(?)");

	    if err != nil {

	        panic(err.Error())

	    }

	    res, err := stmtIns.Exec(entity.Name)

	    if err != nil {

	    	panic(err.Error())

	    }

	    playerId, err := res.LastInsertId()

	    if err != nil {

	        panic(err.Error())

	    }
	    
	    defer stmtIns.Close()

	    entity.Id = int(playerId)

	    return entity.Id;

	}

}

func saveModality(modalities []model.Modality, playerId int , db *sql.DB) {
	
	for _,modality := range modalities {

		modalityId,_ := strconv.Atoi(modality.Id)

		stmtIns, err := db.Prepare("DELETE FROM modality_coverage WHERE id_modality = ? ");

		if err != nil {

	        panic(err.Error())

	    }

	    _, err = stmtIns.Exec(modality.Id)

	    if err != nil {

	    	panic(err.Error())

	    }
	       
	    defer stmtIns.Close()

		if modalityId > 0{

			stmtIns, err := db.Prepare("UPDATE modality SET name=?, price_km=?, time_km=? WHERE id=?");

		    if err != nil {

		        panic(err.Error())

		    }

		    _, err = stmtIns.Exec(modality.Name, modality.PriceKm, modality.TimeKm, modality.Id)

		    if err != nil {

		    	panic(err.Error())

		    }
		       
		    defer stmtIns.Close()

		}else{
			
			stmtIns, err := db.Prepare("INSERT INTO modality (name, price_km, time_km, id_player) VALUES(?, ?, ?, ?)");

		    if err != nil {

		        panic(err.Error())

		    }

		    res, err := stmtIns.Exec(modality.Name, modality.PriceKm, modality.TimeKm, playerId)

		    if err != nil {

		    	panic(err.Error())

		    }

		    id, err := res.LastInsertId()

		    if err != nil {

		        panic(err.Error())

		    }
		    
		    defer stmtIns.Close()

		    modality.Id = strconv.Itoa(int(id))
		}

		saveCoverage(modality, playerId, db)
	}
}

func saveCoverage(modality model.Modality, playerId int , db *sql.DB) {
	
    if len(modality.ModalityCoverage) > 0 {

    	for _,coverage := range modality.ModalityCoverage {

			stmtIns, err := db.Prepare("INSERT INTO modality_coverage (id_modality, zip_code_initial, zip_code_final) VALUES(?, ?, ?)");

		    if err != nil {

		        panic(err.Error())

		    }

		    _, err = stmtIns.Exec(modality.Id, coverage.ZipCodeInitial, coverage.ZipCodeFinal)

		    if err != nil {

		    	panic(err.Error())

		    }
		    
		    defer stmtIns.Close()
		
		}

    }

}

func DeletePlayer(playerId int64){
	
	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err.Error())

	}

	stmtIns, err := db.Prepare("DELETE FROM player where id = ?");

    if err != nil {

        panic(err.Error())

    }

    _, err = stmtIns.Exec(playerId)

    if err != nil {

    	panic(err.Error())

    }
       
    defer stmtIns.Close()

}

func DeleteModality(playerId int64, modality int64){
	
	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err.Error())

	}

	stmtIns, err := db.Prepare("DELETE FROM modality where id = ? and id_player = ?");

    if err != nil {

        panic(err.Error())

    }

    _, err = stmtIns.Exec(modality, playerId)

    if err != nil {

    	panic(err.Error())

    }
       
    defer stmtIns.Close()

}

func FindAllPlayers() (players []model.Player){

	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err)

	}

	rows, err := db.Query("select p.id, p.name, m.id, m.name, m.price_km, m.time_km, mc.zip_code_initial, mc.zip_code_final " + 
							"from player p " +
								"left join modality m on p.id = m.id_player " +
								"left join modality_coverage mc on m.id = mc.id_modality " +
							"where p.id not in (1,2)") 

    if err != nil {

       fmt.Println(err)

        os.Exit(1)

    }

    for rows.Next() {

    	var playerId int
    	
    	var playerName string
    	
    	var modalityId int
    	
    	var modalityName string
    	
    	var modalityPrice float64
    	
    	var modalityTime int
    	
    	var zipCodeInitial string
    	
    	var zipCodeFinal string

        err = rows.Scan(&playerId, &playerName, &modalityId, &modalityName, &modalityPrice, &modalityTime, &zipCodeInitial, &zipCodeFinal)

        if len(players) > 0{

        	var playerExists model.Player

        	for _,player := range players {

        		if player.Id == playerId{

        			playerExists = player

        			var modalityExists model.Modality

        			if len(playerExists.Modalities) > 0 {

        				for _,modality := range player.Modalities {

        					if modality.Id == strconv.Itoa(modalityId) {

        						modalityExists = modality

        						break

        					}

        				}

        				modalityIdInt,_ := strconv.Atoi(modalityExists.Id)

        				if modalityIdInt <= 0 {

        					modalityExists = model.Modality{
        						Id: strconv.Itoa(modalityId), 
        						Name: modalityName, 
        						PriceKm: modalityPrice, 
        						TimeKm: modalityTime, 
        						ModalityCoverage: []model.ModalityCoverage{
        							model.ModalityCoverage{
        								ZipCodeInitial: zipCodeInitial, 
        								ZipCodeFinal: zipCodeFinal,
        							},
        						},
        					}

        					player.Modalities = append(player.Modalities, modalityExists);

        				}else{

        					modalityExists.ModalityCoverage = append(modalityExists.ModalityCoverage, model.ModalityCoverage{ZipCodeInitial: zipCodeInitial, ZipCodeFinal: zipCodeFinal});

        				}

        			}

        		}

        	}

        	if playerExists.Id <= 0{

        		modalityExists := model.Modality{
					Id: strconv.Itoa(modalityId), 
					Name: modalityName, 
					PriceKm: modalityPrice, 
					TimeKm: modalityTime, 
					ModalityCoverage: []model.ModalityCoverage{
						model.ModalityCoverage{
							ZipCodeInitial: zipCodeInitial, 
							ZipCodeFinal: zipCodeFinal,
						},
					},
				}

				players = append(players, model.Player{Id: playerId, Name: playerName, Modalities: []model.Modality{modalityExists}});

        	}

        }else{
        	modalityExists := model.Modality{
				Id: strconv.Itoa(modalityId), 
				Name: modalityName, 
				PriceKm: modalityPrice, 
				TimeKm: modalityTime, 
				ModalityCoverage: []model.ModalityCoverage{
					model.ModalityCoverage{
						ZipCodeInitial: zipCodeInitial, 
						ZipCodeFinal: zipCodeFinal,
					},
				},
			}

			players = append(players, model.Player{Id: playerId, Name: playerName, Modalities: []model.Modality{modalityExists}});
        }

    }
    
    defer db.Close()

    return players

}