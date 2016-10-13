package repository

import( 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"agregador/model"
	"strconv"
)

func SavePlayer(entity model.Player) (model model.Player){

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err.Error())

	}

	defer db.Close()

	savePlayer(entity, db)

	return entity

}

func savePlayer(entity model.Player, db *sql.DB) (id int){
	
	stmtIns, err := db.Prepare("UPDATE player SET name = ?, active = ? where id = ?");

    if err != nil {

        panic(err.Error())

    }

    _, err = stmtIns.Exec(entity.Name, entity.Active, entity.Id)

    if err != nil {

    	panic(err.Error())

    }
       
    defer stmtIns.Close()

    return entity.Id

}

func SaveModality(modality model.Modality) {
	
	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err.Error())

	}

	defer db.Close()

	stmtIns, err := db.Prepare("UPDATE modality SET active=? WHERE id=?");

    if err != nil {

        panic(err.Error())

    }

    _, err = stmtIns.Exec(modality.Active, modality.Id)

    if err != nil {

    	panic(err.Error())

    }
       
    defer stmtIns.Close()

}

func DeletePlayer(playerId int64){
	
	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	
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
	
	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	
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

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err)

	}

	rows, err := db.Query("select p.id, p.name, m.id, m.name, m.price_km, m.time_km, mc.zip_code_initial, mc.zip_code_final, m.price_base, m.price_time, m.minimum_price, p.active, m.active, m.edit_values, p.token  " + 
							"from player p " +
								"left join modality m on p.id = m.id_player " +
								"left join modality_coverage mc on m.id = mc.id_modality ") 

    if err != nil {

       fmt.Println(err)

       return players

    }

    for rows.Next() {

    	var playerId int
    	
    	var playerName string
    	
    	var modalityId int
    	
    	var modalityName string
    	
    	var modalityPrice float64
    	
    	var modalityTime int
    	
    	var zipCodeInitial []byte
    	
    	var zipCodeFinal []byte

    	var priceBase float64
        
        var priceTime float64
        
        var priceMinimum float64

        var activePlayer int

        var activeModality int

        var editValues int

        var token string

        err = rows.Scan(&playerId, &playerName, &modalityId, &modalityName, &modalityPrice, &modalityTime, &zipCodeInitial, &zipCodeFinal, &priceBase, &priceTime, &priceMinimum, &activePlayer, &activeModality, &editValues, &token)

        if err != nil{

        	fmt.Println(err)

       		return players

        }

        var zipCodeFinalAux string

        if zipCodeFinal != nil{

        	zipCodeFinalAux = string(zipCodeFinal) 

        }else{

        	zipCodeFinalAux = "" 

        }

        var zipCodeInitialAux string

        if zipCodeInitial != nil{

        	zipCodeInitialAux = string(zipCodeInitial) 

        }else{

        	zipCodeInitialAux = "" 

        }

        if len(players) > 0{

        	var playerExists model.Player

        	for idx,_ := range players {

        		if players[idx].Id == playerId{

        			playerExists = players[idx]

        			var modalityExists model.Modality

        			if len(playerExists.Modalities) > 0 {

        				for _,modality := range players[idx].Modalities {

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
        						PriceBase: priceBase, 
			                    PriceTime: priceTime, 
			                    PriceMinimum: priceMinimum, 
			                    Active: activeModality,
                				EditValues: editValues,
        						ModalityCoverage: []model.Coverage{
        							model.Coverage{
        								ZipCodeInitial: zipCodeInitialAux, 
        								ZipCodeFinal: zipCodeFinalAux,
        							},
        						},
        					}

        					players[idx].Modalities = append(players[idx].Modalities, modalityExists);

        				}else{

        					modalityExists.ModalityCoverage = append(modalityExists.ModalityCoverage, model.Coverage{ZipCodeInitial: zipCodeInitialAux, ZipCodeFinal: zipCodeFinalAux});

        				}

        			}else{

        				modalityExists = model.Modality{
    						Id: strconv.Itoa(modalityId), 
    						Name: modalityName, 
    						PriceKm: modalityPrice, 
    						TimeKm: modalityTime, 
    						PriceBase: priceBase, 
		                    PriceTime: priceTime, 
		                    PriceMinimum: priceMinimum, 
		                     Active: activeModality,
                			EditValues: editValues,
    						ModalityCoverage: []model.Coverage{
    							model.Coverage{
    								ZipCodeInitial: zipCodeInitialAux, 
    								ZipCodeFinal: zipCodeFinalAux,
    							},
    						},
    					}

        				players[idx].Modalities = append(players[idx].Modalities, modalityExists);

        			}

        		}

        	}

        	if playerExists.Id <= 0{

        		modalityExists := model.Modality{
					Id: strconv.Itoa(modalityId), 
					Name: modalityName, 
					PriceKm: modalityPrice, 
					TimeKm: modalityTime, 
					PriceBase: priceBase, 
                    PriceTime: priceTime, 
                    PriceMinimum: priceMinimum, 
                    Active: activeModality,
                	EditValues: editValues,
					ModalityCoverage: []model.Coverage{
						model.Coverage{
							ZipCodeInitial: zipCodeInitialAux, 
							ZipCodeFinal: zipCodeFinalAux,
						},
					},
				}

				players = append(players, model.Player{Id: playerId, Name: playerName, Active: activePlayer, Token: token, Modalities: []model.Modality{modalityExists}});

        	}

        }else{

        	modalityExists := model.Modality{
				Id: strconv.Itoa(modalityId), 
				Name: modalityName, 
				PriceKm: modalityPrice, 
				TimeKm: modalityTime, 
				PriceBase: priceBase, 
                PriceTime: priceTime, 
                PriceMinimum: priceMinimum, 
                Active: activeModality,
                EditValues: editValues,
				ModalityCoverage: []model.Coverage{
					model.Coverage{
						ZipCodeInitial: zipCodeInitialAux, 
						ZipCodeFinal: zipCodeFinalAux,
					},
				},
			}

			players = append(players, model.Player{Id: playerId, Name: playerName, Active: activePlayer, Token: token, Modalities: []model.Modality{modalityExists}});
			
        }

    }
    
    defer db.Close()

    return players

}