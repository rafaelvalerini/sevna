package repository

import (
	"agregador/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAccessToken() (tokens []model.TokenPlayer) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		panic(err)

	}

	rows, err := db.Query("SELECT id, access_key, id_player FROM player_token where active = 1")

	if err != nil {

		fmt.Println(err)

		panic(err)

	}

	for rows.Next() {

		token := model.TokenPlayer{}

		err = rows.Scan(&token.Id, &token.Token, &token.PlayerId)

		tokens = append(tokens, token)

	}

	defer db.Close()

	return tokens

}
