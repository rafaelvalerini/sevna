package repository

import (
	"agregador/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"os"
)

func FindTokenUber() (token model.TokenUber) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		panic(err)

	}

	rows, err := db.Query("SELECT token_bearer FROM uber_token where active = 1")

	if err != nil {

		fmt.Println(err)

		os.Exit(1)

	}

	var tokens []string

	for rows.Next() {

		var token string

		err = rows.Scan(&token)

		tokens = append(tokens, token)

	}

	token.TokenBearer = getUnicToken(tokens)

	defer db.Close()

	return token

}

func getUnicToken(tokens []string) string {

	index := rand.Intn(len(tokens))

	return tokens[index]

}
