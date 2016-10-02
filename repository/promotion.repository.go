package repository

import( 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"agregador/model"
	"strconv"
)

func SavePromotion(entity model.Promotion, modality int) (model model.Promotion){

	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err.Error())

	}

	defer db.Close()

	id := savePromotion(entity, db)

	saveModalityPromotion(entity, modality, id, db)

	return entity

}

func savePromotion(entity model.Promotion, db *sql.DB) (id int){
	
    idPromotion,_ := strconv.Atoi(entity.Id)

	if idPromotion > 0{

		stmtIns, err := db.Prepare("UPDATE promotion SET name=?, off=?, promotion_code=? where id = ?");

	    if err != nil {

	        panic(err.Error())

	    }

	    _, err = stmtIns.Exec(entity.Name, entity.Off, entity.PromotionCode, idPromotion)

	    if err != nil {

	    	panic(err.Error())

	    }
	       
	    defer stmtIns.Close()

	    return idPromotion

	}else{

		stmtIns, err := db.Prepare("INSERT INTO promotion (name, off, promotion_code) VALUES(?, ?, ?)");

	    if err != nil {

	        panic(err.Error())

	    }

	    res, err := stmtIns.Exec(entity.Name, entity.Off, entity.PromotionCode)

	    if err != nil {

	    	panic(err.Error())

	    }

	    playerId, err := res.LastInsertId()

	    if err != nil {

	        panic(err.Error())

	    }
	    
	    defer stmtIns.Close()

	    entity.Id = strconv.Itoa(int(playerId))

	    return int(playerId);

	}

}

func saveModalityPromotion(entity model.Promotion, modality int, idPromotion int, db *sql.DB) {
	
    modalityId := modality

	stmtIns, err := db.Prepare("INSERT INTO promotion_modality(id_promotion, id_modality, initial_at, final_at, initial_hour, final_hour) VALUES(?, ?, ?, ?, ?, ?)");

    if err != nil {

        panic(err.Error())

    }

    res, err := stmtIns.Exec(idPromotion, modalityId, entity.StartDate, entity.EndDate)

    if err != nil {

    	panic(err.Error())

    }

    id, err := res.LastInsertId()

    if err != nil {

        panic(err.Error())

    }
    
    defer stmtIns.Close()

    promotionModalityId := int(id)

	saveCoveragePromotion(entity, promotionModalityId, modalityId, db)

}

func saveCoveragePromotion(entity model.Promotion, promotionModalityId int , modality int, db *sql.DB) {
	
	stmtIns, err := db.Prepare("DELETE promotion_modality_coverage "+ 
                                " FROM promotion_modality_coverage "+ 
                                " INNER JOIN promotion_modality ON promotion_modality.id = promotion_modality_coverage.id_promotion_modality "+ 
                                " WHERE promotion_modality.id_modality = ? and promotion_modality.id_promotion = ?");

    if err != nil {

        panic(err.Error())

    }

    _, err = stmtIns.Exec(modality, entity.Id)

    if err != nil {

    	panic(err.Error())

    }
    
    defer stmtIns.Close()

    if len(entity.PromotionCoverages) > 0 {

    	for _,coverage := range entity.PromotionCoverages {

			stmtIns, err := db.Prepare("INSERT INTO promotion_modality_coverage (id_promotion_modality, zip_code_initial, zip_code_final) VALUES(?, ?, ?)");

		    if err != nil {

		        panic(err.Error())

		    }

		    _, err = stmtIns.Exec(promotionModalityId, coverage.ZipCodeInitial, coverage.ZipCodeFinal)

		    if err != nil {

		    	panic(err.Error())

		    }
		    
		    defer stmtIns.Close()
		
		}

    }

}

func DeletePromotion(promotion int){
	
	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err.Error())

	}

	stmtIns, err := db.Prepare("DELETE FROM promotion where id = ?");

    if err != nil {

        panic(err.Error())

    }

    _, err = stmtIns.Exec(promotion)

    if err != nil {

    	panic(err.Error())

    }
       
    defer stmtIns.Close()

}

func DeletePromotionModality(promotionModality int){
	
	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    fmt.Println(err)

	}

	stmtIns, err := db.Prepare("DELETE FROM promotion_modality where id = ?");

    if err != nil {

        fmt.Println(err)

    }

    _, err = stmtIns.Exec(promotionModality)

    if err != nil {

    	fmt.Println(err)

    }
       
    defer stmtIns.Close()

}

func FindPromotion(player int, modality int) (promotions []model.Promotion){

	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    fmt.Println(err)

        defer db.Close()

        return promotions

	}

	rows, err := db.Query("select p.id, p.name, p.off, p.promotion_code, pmc.zip_code_initial, pmc.zip_code_final, "+
                            "pm.initial_at, pm.final_at, pm.initial_hour, pm.final_hour " + 
						"from promotion p " + 
							"inner join promotion_modality pm on p.id = pm.id_promotion " + 
							"left join promotion_modality_coverage pmc on pmc.id_promotion_modality = pm.id " + 
						"where pm.id_modality = ?", modality) 



    if err != nil {

        fmt.Println(err)

        defer db.Close()

        return promotions

    }

    for rows.Next() {

    	var promotionId int
    	
    	var promotionName string
    	
    	var off float64
    	
    	var zipCodeInitial string
    	
    	var zipCodeFinal string

        var initialAt []byte

        var finalAt []byte

        err = rows.Scan(&promotionId, &promotionName, &off, &zipCodeInitial, &zipCodeFinal, &initialAt, &finalAt)

        if len(promotions) > 0{

        	var promotionExists model.Promotion

        	for idx,_ := range promotions {

                promotionBaseId,_ := strconv.Atoi(promotions[idx].Id)

        		if promotionBaseId == promotionId{

        			promotionExists = promotions[idx]

        			promotions[idx].PromotionCoverages = append(promotions[idx].PromotionCoverages, model.Coverage{ZipCodeInitial: zipCodeInitial, ZipCodeFinal: zipCodeFinal,});

        		}

        	}

        	if promotionExists.Id == ""{

        		promotionExists = model.Promotion{
                    Id: strconv.Itoa(promotionId), 
                    Name: promotionName, 
                    Off: off, 
                    StartDate: string(initialAt),
                    EndDate: string(finalAt),
                    PromotionCoverages: []model.Coverage{
                        model.Coverage{
                            ZipCodeInitial: zipCodeInitial, 
                            ZipCodeFinal: zipCodeFinal,
                        },
                    },
                }

				promotions = append(promotions, promotionExists);

        	}

        }else{

        	promotionExists := model.Promotion{
                Id: strconv.Itoa(promotionId), 
                Name: promotionName, 
                Off: off, 
                StartDate: string(initialAt),
                EndDate: string(finalAt),
                PromotionCoverages: []model.Coverage{
                    model.Coverage{
                        ZipCodeInitial: zipCodeInitial, 
                        ZipCodeFinal: zipCodeFinal,
                    },
                },
            }

            promotions = append(promotions, promotionExists);
			
        }

    }
    
    defer db.Close()

    return promotions

}

func FindAllPromotions() (promotions []model.Promotion){

    db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
    
    if err != nil {

        fmt.Println(err)

        defer db.Close()

        return promotions

    }

    rows, err := db.Query("select p.id, m.name, p.name, p.off, p.limit_off, p.new_modality, p.text, "+
                            "pa.monday, pa.tuesday, pa.wednesday, pa.thursday, pa.friday, pa.saturday, pa.sunday, "+
                            "pa.start_hour, pa.end_hour, "+
                            "pmc.zip_code_initial, pmc.zip_code_final, pmc.state, pmc.city "+
                        "from promotion p  "+
                            "left join promotion_available pa on p.id = pa.id_promotion "+
                            "inner join promotion_modality pm on p.id = pm.id_promotion "+
                            "inner join modality m on pm.id_modality = m.id "+
                            "left join promotion_modality_coverage pmc on pm.id = pmc.id_promotion_modality  "+
                        "where p.active = 1 and (pm.initial_at is null or NOW() between pm.initial_at and pm.final_at)") 


    if err != nil {

        fmt.Println(err)

        defer db.Close()

        return promotions

    }

    for rows.Next() {

        var promotionId int

        var modalityName string

        var promotionName string
        
        var off float64

        var limitOff float64

        var newModality int
        
        var text string

        var monday int

        var tuesday int

        var wednesday int

        var thursday int

        var friday int

        var saturday int

        var sunday int

        var endHour []byte

        var startHour []byte

        var zipCodeInitial []byte
        
        var zipCodeFinal []byte

        var state []byte

        var city []byte
        
         err = rows.Scan(&promotionId, &modalityName, &promotionName, &off, &limitOff, &newModality, &text, &monday, &tuesday, &wednesday, &thursday, &friday, &saturday, &sunday, &startHour, &endHour, &zipCodeInitial, &zipCodeFinal, &state, &city)

        if err != nil{

            fmt.Println(err)

            defer db.Close()

            return promotions

        }

        if len(promotions) > 0{

            var promotionExists model.Promotion

            for idx,_ := range promotions {

                promotionBaseId,_ := strconv.Atoi(promotions[idx].Id)

                if promotionBaseId == promotionId && modalityName == promotions[idx].Modality{

                    promotionExists = promotions[idx]

                    if zipCodeInitial != nil || zipCodeFinal != nil || city != nil || state != nil{

                        add := true

                        for _,cov := range promotions[idx].PromotionCoverages {

                            if cov.ZipCodeInitial == string(zipCodeInitial) && cov.ZipCodeFinal == string(zipCodeFinal) && cov.City == string(city) && cov.State == string(state){
                                add = false
                            }
                            
                        }

                        if add {

                            promotions[idx].PromotionCoverages = append(promotions[idx].PromotionCoverages, model.Coverage{
                                ZipCodeInitial: string(zipCodeInitial), 
                                ZipCodeFinal: string(zipCodeFinal), 
                                City: string(city), 
                                State: string(state),
                            })    

                        }

                    }

                    promotions[idx].PromotionAvailable = append(promotions[idx].PromotionAvailable, model.Available{
                        Monday: monday, 
                        Tuesday: tuesday, 
                        Wednesday: wednesday, 
                        Thursday: thursday, 
                        Friday: friday, 
                        Saturday: saturday,
                        Sunday: sunday, 
                        StartHour:  
                        string(startHour), 
                        EndHour: string(endHour),
                    })

                }

            }

            if promotionExists.Id == ""{

                promotionExists = model.Promotion{
                    Id: strconv.Itoa(promotionId), 
                    Name: promotionName, 
                    Off: off, 
                    Modality: modalityName,
                    NewModality: newModality,
                    PromotionAvailable: []model.Available{
                        model.Available{
                            Monday: monday,
                            Tuesday: tuesday,
                            Wednesday: wednesday,
                            Thursday: thursday,
                            Friday: friday,
                            Saturday: saturday,
                            Sunday: sunday,
                            StartHour:  string(startHour),
                            EndHour: string(endHour),
                        },
                    },
                    PromotionCoverages: []model.Coverage{
                        model.Coverage{
                            ZipCodeInitial: string(zipCodeInitial), 
                            ZipCodeFinal: string(zipCodeFinal),
                            State: string(state),
                            City: string(city),
                        },
                    },
                }

                promotions = append(promotions, promotionExists);

            }

        }else{

            promotionExists := model.Promotion{
                Id: strconv.Itoa(promotionId), 
                Name: promotionName, 
                Off: off, 
                Modality: modalityName,
                NewModality: newModality,
                PromotionAvailable: []model.Available{
                    model.Available{
                        Monday: monday,
                        Tuesday: tuesday,
                        Wednesday: wednesday,
                        Thursday: thursday,
                        Friday: friday,
                        Saturday: saturday,
                        Sunday: sunday,
                        StartHour:  string(startHour),
                        EndHour: string(endHour),
                    },
                },
                PromotionCoverages: []model.Coverage{
                    model.Coverage{
                        ZipCodeInitial: string(zipCodeInitial), 
                        ZipCodeFinal: string(zipCodeFinal),
                        State: string(state),
                        City: string(city),
                    },
                },
            }

            promotions = append(promotions, promotionExists);
            
        }

    }
    
    defer db.Close()

    return promotions

}