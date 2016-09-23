package repository

import( 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"agregador/model"
    "strconv"
)

func SearchPlayersDefault(position model.Position) (players []model.Player){

	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err)

	}

	rows, err := db.Query("select " + 
                            "p.id, " + 
                            "p.name, " + 
                            "m.id, " + 
                            "m.name, " + 
                            "m.price_km, " + 
                            "m.time_km, " + 
                            "promo.id, " + 
                            "promo.name, " + 
                            "promo.off, " + 
                            "promo.promotion_code, " + 
                            "promo.initial_at, " + 
                            "promo.final_at, " + 
                            "promo.initial_hour, " + 
                            "promo.final_hour " + 
                        "from " + 
                            "player p  " + 
                            "inner join modality m on p.id = m.id_player  " + 
                            "inner join modality_coverage mc on mc.id_modality = m.id   " + 
                            "left join ( " + 
                                "select promo.id as id,  " + 
                                    "promo.name as name,  " + 
                                    "promo.off as off,  " + 
                                    "promo.promotion_code as promotion_code, " + 
                                    "pm.id_modality as id_modality, " +
                                    "pm.initial_at as initial_at, " +
                                    "pm.final_at as final_at, " +
                                    "pm.initial_hour as initial_hour, " +
                                    "pm.final_hour as final_hour " + 
                                "from promotion_modality pm  " + 
                                    "left join promotion promo on promo.id = pm.id_promotion  " + 
                                    "left join promotion_modality_coverage pmc on pmc.id_promotion_modality = pm.id " + 
                                "where  " + 
                                    "promo.id is null or pmc.id is null or (? between LEFT(pmc.zip_code_initial, 5) and LEFT(pmc.zip_code_final, 5)) " + 
                                    ") promo on promo.id_modality = m.id " + 
                        "where " + 
                            "? between LEFT(mc.zip_code_initial, 5) and LEFT(mc.zip_code_final, 5)", position.ZipCode, position.ZipCode) 

    if err != nil {

        panic(err);

    }

    for rows.Next() {
        var idPlayer int
        var namePlayer string
        var idModality int
        var nameModality string
        var priceKm float64
        var timeKm int
        var idPromo string
        var namePromo string
        var off float64
        var promoCode string
        var initialAt int64
        var finalAt int64
        var initialHour string
        var finalHour string
        err = rows.Scan(&idPlayer, &namePlayer, &idModality, &nameModality, &priceKm, &timeKm, &idPromo, &namePromo, &off, &promoCode, &initialAt, &finalAt, &initialHour, &finalHour)

        element := model.Player{
                Id: idPlayer, 
                Name: namePlayer, 
                Modality: model.Modality{
                    Id: strconv.Itoa(idModality), 
                    Name: nameModality, 
                    PriceKm: priceKm, 
                    TimeKm: timeKm,
                    Promotion: model.Promotion{
                        Id: idPromo,
                        Name: namePromo,
                        Off: off,
                        PromotionCode: promoCode,
                        StartDate: initialAt,
                        EndDate: finalAt,
                        StartHour: initialHour,
                        EndHour: finalHour,
                    },
                },
            }

        players = append(players, element)
        
    }
    
    defer db.Close()

    return players

}

