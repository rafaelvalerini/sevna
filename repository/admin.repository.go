package repository

import (
	"agregador/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"time"
)

func CountEstimates() (cont int64) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")

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

func CountModalities() (modalities []model.MoreUser) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")

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
			Value:    count,
		}

		modalities = append(modalities, element)

	}

	defer db.Close()

	return modalities

}

func SearchByNotification(state string, city string, initial int) (notifications []string) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		panic(err)

	}

	rows, err := db.Query(" select distinct device.id_notification "+
		" from config_device device  "+
		" inner join search s on s.id = device.id_search  "+
		" inner join search_address address on (address.id = s.start_address_id or address.id = s.end_address_id)  "+
		" where ( = '' or ? = address.state) and (? = '' or ? = address.city)  "+
		" order by device.id_notification  "+
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

func SumSavedEstimates() (model model.MetaFloat64) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		panic(err)

	}

	rows, err := db.Query("select " +
		"sum( tot.soma ) " +
		"from " +
		"( " +
		"select " +
		"value_max.id, " +
		"value_max.count_max, " +
		"sel.count_min, " +
		"( " +
		"value_max.count_max - sel.count_min " +
		") as soma " +
		"from " +
		"( " +
		"select " +
		"s.id as id, " +
		"max( cast( SUBSTRING_INDEX( REPLACE( REPLACE( REPLACE( REPLACE( REPLACE( REPLACE( sr.tax_value, 'N/A', '10' ), 'Unavailable', '10' ), '€', '' ), ' ', '' ), ',', '.' ), 'R$', '' ), '-', 1 ) as decimal( 8, 2 ) ) ) as count_max " +
		"from " +
		"search s inner join search_results sr on " +
		"s.id = sr.id_search " +
		"group by " +
		"s.id " +
		") value_max left join( " +
		"select " +
		"sr.id_search as id, " +
		"min( cast( SUBSTRING_INDEX( REPLACE( REPLACE( REPLACE( REPLACE( REPLACE( REPLACE( sr.tax_value, 'N/A', '10' ), 'Unavailable', '10' ), '€', '' ), ' ', '' ), ',', '.' ), 'R$', '' ), '-', 1 ) as decimal( 8, 2 ) ) ) as count_min " +
		"from " +
		"search_results sr inner join search_selected ss on " +
		"sr.id = ss.id_search_results " +
		" group by " +
		"sr.id_search " +
		") sel on " +
		"value_max.id = sel.id " +
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

func FindAnalytics(state string, city string, player string, modality string, startAt string, endAt string) (result []model.ResultAnalytics) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		panic(err)

	}

	defer db.Close()

	layout := "2006-01-02T15:04:05.000Z"

	start_date, err := time.Parse(layout, startAt)

	end_date, err := time.Parse(layout, endAt)

	query := fmt.Sprintf("select "+
		"distinct  "+
		"startAddress.date_time as DateTime, "+
		"CASE "+
		"WHEN selected.type_open IS NULL "+
		"THEN 0 "+
		"ELSE selected.type_open "+
		"END as TypeOpen,"+
		"startAddress.address as StartAddress, "+
		"endAddress.address as EndAddress, "+
		"p.name as Player, "+
		"selected.id_modality as Modality, "+
		"selected.tax_value as Value, "+
		"selected.promotion as Promotion, "+
		"c.operation_system as OperationSystem, "+
		"startAddress.city as StartCity, "+
		"startAddress.state as StartState, "+
		"endAddress.city as EndCity, "+
		"endAddress.state as EndState, "+
		"c.operation_system_version as OperationSystemVersion, "+
		"c.type_connection as TypeConnection, "+

		"c.version_app as VersionApp "+
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
		"res.tax_value as tax_value, "+
		"se.promotion as promotion, "+
		"se.type_open as type_open "+
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

	if state != "" {
		query = query + " and startAddress.state = '" + state + "'"
	}

	if city != "" {
		query = query + " and startAddress.city = '" + city + "'"
	}

	if player != "" {
		playerId, _ := strconv.Atoi(player)
		query = query + " and p.id = " + fmt.Sprintf("%d", playerId)
	}

	if modality != "" {
		query = query + " and UPPER(selected.id_modality) = '" + modality + "'"
	}

	query = query + " order by startAddress.date_time desc "

	rows, err := db.Query(query)

	if err != nil {

		fmt.Println(err)

	}

	defer rows.Close()

	for rows.Next() {

		var res model.ResultAnalytics

		var playerResult sql.NullString

		var modalityResult sql.NullString

		var valueResult sql.NullString

		var promotion sql.NullString

		var version sql.NullString

		var typeConnection sql.NullString

		err2 := rows.Scan(&res.DateTime, &res.TypeOpen, &res.StartAddress, &res.EndAddress, &playerResult, &modalityResult, &valueResult, &promotion, &res.OperationSystem, &res.StartCity, &res.StartState, &res.EndCity, &res.EndState, &res.OperationSystemVersion, &typeConnection, &version)

		if err2 != nil {
			fmt.Println(err2)
			continue
		}

		if typeConnection.Valid {
			res.TypeConnection = typeConnection.String
		}

		if playerResult.Valid {
			res.Player = playerResult.String
		}

		if modalityResult.Valid {
			res.Modality = modalityResult.String
		}

		if valueResult.Valid {
			res.Value = valueResult.String
		}

		if promotion.Valid {
			res.Promotion = promotion.String
		}

		if version.Valid {
			res.VersionApp = version.String
		}

		result = append(result, res)

	}

	return result

}

func CountPromotions() (modalities []model.MoreUser) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		panic(err)

	}

	rows, err := db.Query("select count(s.id) as count, s.promotion as promotion from search_selected s where s.promotion is not null group by s.promotion")

	if err != nil {

		fmt.Println(err)

		os.Exit(1)

	}

	for rows.Next() {

		var promotion string

		var count int64

		err = rows.Scan(&count, &promotion)

		element := model.MoreUser{
			Modality: promotion,
			Value:    count,
		}

		modalities = append(modalities, element)

	}

	defer db.Close()

	return modalities

}
