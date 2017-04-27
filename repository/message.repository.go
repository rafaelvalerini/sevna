package repository

import (
	"agregador/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
)

func FindMessagesStart(lastRecord int) (message model.MessageStart) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()

	if err != nil {

		panic(err)

	}

	rows, err := db.Query("select id, message from message_start where NOW() between start_at and end_at AND id > " + strconv.Itoa(lastRecord) + " order by id desc LIMIT 1")

	if err != nil {

		fmt.Println(err)

		os.Exit(1)

	}

	for rows.Next() {

		var messageNullable sql.NullString

		err = rows.Scan(&message.Id, &messageNullable)

		if messageNullable.Valid {
			message.Message = messageNullable.String
		}

		return message

	}

	return message

}
