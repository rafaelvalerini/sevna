package repository

import( 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
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

    res, err := stmtIns.Exec(idPromotion, modalityId, entity.StartDate, entity.EndDate, entity.StartHour, entity.EndHour)

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

	    panic(err.Error())

	}

	stmtIns, err := db.Prepare("DELETE FROM promotion_modality where id = ?");

    if err != nil {

        panic(err.Error())

    }

    _, err = stmtIns.Exec(promotionModality)

    if err != nil {

    	panic(err.Error())

    }
       
    defer stmtIns.Close()

}

func FindPromotion(player int, modality int) (promotions []model.Promotion){

	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err)

	}

	rows, err := db.Query("select p.id, p.name, p.off, p.promotion_code, pmc.zip_code_initial, pmc.zip_code_final, "+
                            "pm.initial_at, pm.final_at, pm.initial_hour, pm.final_hour " + 
						"from promotion p " + 
							"inner join promotion_modality pm on p.id = pm.id_promotion " + 
							"left join promotion_modality_coverage pmc on pmc.id_promotion_modality = pm.id " + 
						"where pm.id_modality = ?", modality) 



    if err != nil {

       fmt.Println(err)

        os.Exit(1)

    }

    for rows.Next() {

    	var promotionId int
    	
    	var promotionName string
    	
    	var off float64
    	
    	var promotionCode string
    	
    	var zipCodeInitial string
    	
    	var zipCodeFinal string

        var initialAt int64

        var finalAt int64

        var initialHour string

        var finalHour string
    	
        err = rows.Scan(&promotionId, &promotionName, &off, &promotionCode, &zipCodeInitial, &zipCodeFinal, &initialAt, &finalAt, &initialHour, &finalHour)

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
                    PromotionCode: promotionCode,
                    StartDate: initialAt,
                    EndDate: finalAt,
                    StartHour: initialHour,
                    EndHour: finalHour,
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
                PromotionCode: promotionCode,
                StartDate: initialAt,
                EndDate: finalAt,
                StartHour: initialHour,
                EndHour: finalHour,
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