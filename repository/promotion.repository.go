package repository

import (
	"agregador/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func SavePromotion(entity model.Promotion) (model model.Promotion) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		panic(err.Error())

	}

	defer db.Close()

	if entity.Id > 0 {

		DeletePromotion(entity.Id)

		entity.Id = 0

	}

	id := savePromotion(entity, db)

	saveModalityPromotion(entity, id, db)

	savePromotionAvailabled(entity, id, db)

	return entity

}

func savePromotion(entity model.Promotion, db *sql.DB) (id int) {

	stmtIns, err := db.Prepare("INSERT INTO promotion (name, off, limit_off, new_modality, active, text) VALUES(?, ?, ?, 1, 1, ?)")

	if err != nil {

		panic(err.Error())

	}

	res, err := stmtIns.Exec(entity.Name, entity.Off, entity.LimitOff, entity.Description)

	if err != nil {

		panic(err.Error())

	}

	playerId, err := res.LastInsertId()

	if err != nil {

		panic(err.Error())

	}

	defer stmtIns.Close()

	entity.Id = int(playerId)

	return int(playerId)

}

func saveModalityPromotion(entity model.Promotion, idPromotion int, db *sql.DB) {

	for i := 0; i < len(entity.Modalities); i++ {

		stmtIns, err := db.Prepare("INSERT INTO promotion_modality(id_promotion, id_modality, initial_at, final_at, exibition_name) VALUES(?, ?, ?, ?, ?)")

		if err != nil {

			panic(err.Error())

		}

		res, err := stmtIns.Exec(idPromotion, entity.Modalities[i].Id, entity.StartDate, entity.EndDate, entity.ExibitionName)

		if err != nil {

			panic(err.Error())

		}

		id, err := res.LastInsertId()

		if err != nil {

			panic(err.Error())

		}

		defer stmtIns.Close()

		promotionModalityId := int(id)

		saveCoveragePromotion(entity, promotionModalityId, db)

	}

}

func savePromotionAvailabled(entity model.Promotion, idPromotion int, db *sql.DB) {

	for i := 0; i < len(entity.PromotionAvailable); i++ {

		stmtIns, err := db.Prepare("INSERT INTO vah.promotion_available (monday, tuesday, wednesday, thursday, friday, saturday, sunday, start_hour, end_hour, id_promotion) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")

		if err != nil {

			panic(err.Error())

		}

		_, err = stmtIns.Exec(entity.PromotionAvailable[i].Monday, entity.PromotionAvailable[i].Tuesday, entity.PromotionAvailable[i].Wednesday, entity.PromotionAvailable[i].Thursday, entity.PromotionAvailable[i].Friday, entity.PromotionAvailable[i].Saturday, entity.PromotionAvailable[i].Sunday, entity.PromotionAvailable[i].StartHour, entity.PromotionAvailable[i].EndHour, idPromotion)

		if err != nil {

			panic(err.Error())

		}

		defer stmtIns.Close()

	}

}

func saveCoveragePromotion(entity model.Promotion, promotionModalityId int, db *sql.DB) {

	if len(entity.PromotionCoverages) > 0 {

		for _, coverage := range entity.PromotionCoverages {

			stmtIns, err := db.Prepare("INSERT INTO promotion_modality_coverage (id_promotion_modality, state, city) VALUES(?, ?, ?)")

			if err != nil {

				panic(err.Error())

			}

			_, err = stmtIns.Exec(promotionModalityId, coverage.State, coverage.City)

			if err != nil {

				panic(err.Error())

			}

			defer stmtIns.Close()

		}

	}

}

func DeletePromotion(promotion int) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		panic(err.Error())

	}

	stmtIns, err := db.Prepare("DELETE FROM promotion where id = ?")

	if err != nil {

		panic(err.Error())

	}

	_, err = stmtIns.Exec(promotion)

	if err != nil {

		panic(err.Error())

	}

	defer stmtIns.Close()

}

func DeletePromotionModality(promotionModality int) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		fmt.Println(err)

	}

	stmtIns, err := db.Prepare("DELETE FROM promotion_modality where id = ?")

	if err != nil {

		fmt.Println(err)

	}

	_, err = stmtIns.Exec(promotionModality)

	if err != nil {

		fmt.Println(err)

	}

	defer stmtIns.Close()

}

func FindPromotion(player int, modality int) (promotions []model.Promotion) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		fmt.Println(err)

		defer db.Close()

		return promotions

	}

	rows, err := db.Query("select p.id, p.name, p.off, p.promotion_code, pmc.zip_code_initial, pmc.zip_code_final, "+
		"pm.initial_at, pm.final_at, pm.initial_hour, pm.final_hour "+
		"from promotion p "+
		"inner join promotion_modality pm on p.id = pm.id_promotion "+
		"left join promotion_modality_coverage pmc on pmc.id_promotion_modality = pm.id "+
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

		if len(promotions) > 0 {

			var promotionExists model.Promotion

			for idx, _ := range promotions {

				promotionBaseId := promotions[idx].Id

				if promotionBaseId == promotionId {

					promotionExists = promotions[idx]

					promotions[idx].PromotionCoverages = append(promotions[idx].PromotionCoverages, model.Coverage{ZipCodeInitial: zipCodeInitial, ZipCodeFinal: zipCodeFinal})

				}

			}

			if promotionExists.Id == 0 {

				promotionExists = model.Promotion{
					Id:        promotionId,
					Name:      promotionName,
					Off:       off,
					StartDate: string(initialAt),
					EndDate:   string(finalAt),
					PromotionCoverages: []model.Coverage{
						model.Coverage{
							ZipCodeInitial: zipCodeInitial,
							ZipCodeFinal:   zipCodeFinal,
						},
					},
				}

				promotions = append(promotions, promotionExists)

			}

		} else {

			promotionExists := model.Promotion{
				Id:        promotionId,
				Name:      promotionName,
				Off:       off,
				StartDate: string(initialAt),
				EndDate:   string(finalAt),
				PromotionCoverages: []model.Coverage{
					model.Coverage{
						ZipCodeInitial: zipCodeInitial,
						ZipCodeFinal:   zipCodeFinal,
					},
				},
			}

			promotions = append(promotions, promotionExists)

		}

	}

	defer db.Close()

	return promotions

}

func FindAllPromotions() (promotions []model.Promotion) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		fmt.Println(err)

		defer db.Close()

		return promotions

	}

	rows, err := db.Query("select p.id, m.name, p.name, p.off, p.limit_off, p.new_modality, p.text, pm.exibition_name, " +
		"pa.monday, pa.tuesday, pa.wednesday, pa.thursday, pa.friday, pa.saturday, pa.sunday, " +
		"pa.start_hour, pa.end_hour, " +
		"pmc.zip_code_initial, pmc.zip_code_final, pmc.state, pmc.city " +
		"from promotion p  " +
		"left join promotion_available pa on p.id = pa.id_promotion " +
		"inner join promotion_modality pm on p.id = pm.id_promotion " +
		"inner join modality m on pm.id_modality = m.id " +
		"left join promotion_modality_coverage pmc on pm.id = pmc.id_promotion_modality  " +
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

		var exibitionName string

		err = rows.Scan(&promotionId, &modalityName, &promotionName, &off, &limitOff, &newModality, &text, &exibitionName, &monday, &tuesday, &wednesday, &thursday, &friday, &saturday, &sunday, &startHour, &endHour, &zipCodeInitial, &zipCodeFinal, &state, &city)

		if err != nil {

			fmt.Println(err)

			defer db.Close()

			return promotions

		}

		if len(promotions) > 0 {

			var promotionExists model.Promotion

			for idx, _ := range promotions {

				if promotions[idx].Id == promotionId && modalityName == promotions[idx].Modality {

					promotionExists = promotions[idx]

					if zipCodeInitial != nil || zipCodeFinal != nil || city != nil || state != nil {

						add := true

						for _, cov := range promotions[idx].PromotionCoverages {

							if cov.ZipCodeInitial == string(zipCodeInitial) && cov.ZipCodeFinal == string(zipCodeFinal) && cov.City == string(city) && cov.State == string(state) {
								add = false
							}

						}

						if add {

							promotions[idx].PromotionCoverages = append(promotions[idx].PromotionCoverages, model.Coverage{
								ZipCodeInitial: string(zipCodeInitial),
								ZipCodeFinal:   string(zipCodeFinal),
								City:           string(city),
								State:          string(state),
							})

						}

					}

					promotions[idx].PromotionAvailable = append(promotions[idx].PromotionAvailable, model.Available{
						Monday:    monday,
						Tuesday:   tuesday,
						Wednesday: wednesday,
						Thursday:  thursday,
						Friday:    friday,
						Saturday:  saturday,
						Sunday:    sunday,
						StartHour: string(startHour),
						EndHour:   string(endHour),
					})

				}

			}

			if promotionExists.Id == 0 {

				promotionExists = model.Promotion{
					Id:            promotionId,
					Name:          promotionName,
					Off:           off,
					Modality:      modalityName,
					NewModality:   newModality,
					ExibitionName: exibitionName,
					Description:   text,
					LimitOff:      limitOff,
					PromotionAvailable: []model.Available{
						model.Available{
							Monday:    monday,
							Tuesday:   tuesday,
							Wednesday: wednesday,
							Thursday:  thursday,
							Friday:    friday,
							Saturday:  saturday,
							Sunday:    sunday,
							StartHour: string(startHour),
							EndHour:   string(endHour),
						},
					},
					PromotionCoverages: []model.Coverage{
						model.Coverage{
							ZipCodeInitial: string(zipCodeInitial),
							ZipCodeFinal:   string(zipCodeFinal),
							State:          string(state),
							City:           string(city),
						},
					},
				}

				promotions = append(promotions, promotionExists)

			}

		} else {

			promotionExists := model.Promotion{
				Id:            promotionId,
				Name:          promotionName,
				Off:           off,
				Modality:      modalityName,
				NewModality:   newModality,
				ExibitionName: exibitionName,
				Description:   text,
				LimitOff:      limitOff,
				PromotionAvailable: []model.Available{
					model.Available{
						Monday:    monday,
						Tuesday:   tuesday,
						Wednesday: wednesday,
						Thursday:  thursday,
						Friday:    friday,
						Saturday:  saturday,
						Sunday:    sunday,
						StartHour: string(startHour),
						EndHour:   string(endHour),
					},
				},
				PromotionCoverages: []model.Coverage{
					model.Coverage{
						ZipCodeInitial: string(zipCodeInitial),
						ZipCodeFinal:   string(zipCodeFinal),
						State:          string(state),
						City:           string(city),
					},
				},
			}

			promotions = append(promotions, promotionExists)

		}

	}

	defer db.Close()

	return promotions

}

func FindAllPromotionsGroupModality() (promotions []model.Promotion) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		fmt.Println(err)

		defer db.Close()

		return promotions

	}

	rows, err := db.Query("select p.id, m.name, p.name, p.off, p.limit_off, p.new_modality, p.text, pm.exibition_name, " +
		"pa.monday, pa.tuesday, pa.wednesday, pa.thursday, pa.friday, pa.saturday, pa.sunday, " +
		"pa.start_hour, pa.end_hour, " +
		"pmc.zip_code_initial, pmc.zip_code_final, pmc.state, pmc.city, m.id, p.active, pm.initial_at, pm.final_at, m.id_player " +
		"from promotion p  " +
		"left join promotion_available pa on p.id = pa.id_promotion " +
		"inner join promotion_modality pm on p.id = pm.id_promotion " +
		"inner join modality m on pm.id_modality = m.id " +
		"left join promotion_modality_coverage pmc on pm.id = pmc.id_promotion_modality " +
		" order by p.name ")

	if err != nil {

		fmt.Println(err)

		defer db.Close()

		return promotions

	}

	for rows.Next() {

		var promotionId int

		var modalityName string

		var modalityId int

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

		var exibitionName string

		var active int

		var startDate []byte

		var endDate []byte

		var player int

		err = rows.Scan(&promotionId, &modalityName, &promotionName, &off, &limitOff, &newModality, &text, &exibitionName, &monday, &tuesday, &wednesday, &thursday, &friday, &saturday, &sunday, &startHour, &endHour, &zipCodeInitial, &zipCodeFinal, &state, &city, &modalityId, &active, &startDate, &endDate, &player)

		if err != nil {

			fmt.Println(err)

			defer db.Close()

			return promotions

		}

		if len(promotions) > 0 {

			var promotionExists model.Promotion

		promoFor:
			for idx, _ := range promotions {

				if promotions[idx].Id == promotionId {

					promotionExists = promotions[idx]

					if zipCodeInitial != nil || zipCodeFinal != nil || city != nil || state != nil {

						add := true

						for _, cov := range promotions[idx].PromotionCoverages {

							if cov.ZipCodeInitial == string(zipCodeInitial) && cov.ZipCodeFinal == string(zipCodeFinal) && cov.City == string(city) && cov.State == string(state) {
								add = false
							}

						}

						if add {

							promotions[idx].PromotionCoverages = append(promotions[idx].PromotionCoverages, model.Coverage{
								ZipCodeInitial: string(zipCodeInitial),
								ZipCodeFinal:   string(zipCodeFinal),
								City:           string(city),
								State:          string(state),
							})

						}

					}

					add := true

					for _, cov := range promotions[idx].PromotionAvailable {

						if cov.Monday == monday && cov.Tuesday == tuesday && cov.Wednesday == wednesday && cov.Thursday == thursday && cov.Friday == friday && cov.Saturday == saturday && cov.Sunday == sunday && cov.StartHour == string(startHour) && cov.EndHour == string(endHour) {
							add = false
						}

					}

					if add {

						promotions[idx].PromotionAvailable = append(promotions[idx].PromotionAvailable, model.Available{
							Monday:    monday,
							Tuesday:   tuesday,
							Wednesday: wednesday,
							Thursday:  thursday,
							Friday:    friday,
							Saturday:  saturday,
							Sunday:    sunday,
							StartHour: string(startHour),
							EndHour:   string(endHour),
						})

					}

					for _, modal := range promotions[idx].Modalities {

						if modal.Id == strconv.Itoa(modalityId) {

							continue promoFor

						}

					}

					promotions[idx].Modalities = append(promotions[idx].Modalities, model.ModalitySimple{Id: strconv.Itoa(modalityId), Name: modalityName, Player: player})

				}

			}

			if promotionExists.Id == 0 {

				promotionExists = model.Promotion{
					Id:        promotionId,
					Active:    active,
					Name:      promotionName,
					StartDate: string(startDate),
					EndDate:   string(endDate),
					Off:       off,
					Modalities: []model.ModalitySimple{
						model.ModalitySimple{
							Id:     strconv.Itoa(modalityId),
							Name:   modalityName,
							Player: player,
						},
					},
					NewModality:   newModality,
					ExibitionName: exibitionName,
					Description:   text,
					LimitOff:      limitOff,
					PromotionAvailable: []model.Available{
						model.Available{
							Monday:    monday,
							Tuesday:   tuesday,
							Wednesday: wednesday,
							Thursday:  thursday,
							Friday:    friday,
							Saturday:  saturday,
							Sunday:    sunday,
							StartHour: string(startHour),
							EndHour:   string(endHour),
						},
					},
					PromotionCoverages: []model.Coverage{
						model.Coverage{
							ZipCodeInitial: string(zipCodeInitial),
							ZipCodeFinal:   string(zipCodeFinal),
							State:          string(state),
							City:           string(city),
						},
					},
				}

				promotions = append(promotions, promotionExists)

			}

		} else {

			promotionExists := model.Promotion{
				Id:        promotionId,
				Active:    active,
				Name:      promotionName,
				Off:       off,
				StartDate: string(startDate),
				EndDate:   string(endDate),
				Modalities: []model.ModalitySimple{
					model.ModalitySimple{
						Id:     strconv.Itoa(modalityId),
						Name:   modalityName,
						Player: player,
					},
				},
				NewModality:   newModality,
				ExibitionName: exibitionName,
				Description:   text,
				LimitOff:      limitOff,
				PromotionAvailable: []model.Available{
					model.Available{
						Monday:    monday,
						Tuesday:   tuesday,
						Wednesday: wednesday,
						Thursday:  thursday,
						Friday:    friday,
						Saturday:  saturday,
						Sunday:    sunday,
						StartHour: string(startHour),
						EndHour:   string(endHour),
					},
				},
				PromotionCoverages: []model.Coverage{
					model.Coverage{
						ZipCodeInitial: string(zipCodeInitial),
						ZipCodeFinal:   string(zipCodeFinal),
						State:          string(state),
						City:           string(city),
					},
				},
			}

			promotions = append(promotions, promotionExists)

		}

	}

	defer db.Close()

	return promotions

}
