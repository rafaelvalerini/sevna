package repository

import (
	"agregador/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
	"strings"
)

func Login(mail string, password string) (user model.User) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		panic(err)

	}

	rows, err := db.Query("SELECT id, mail, username, token FROM user where mail = ? and password = ?", mail, password)

	if err != nil {

		fmt.Println(err)

		os.Exit(1)

	}

	for rows.Next() {

		var id int64
		var username string
		var mail string
		var token string

		err = rows.Scan(&id, &mail, &username, &token)

		user = model.User{
			Id:       id,
			Username: username,
			Mail:     mail,
			Token:    token,
		}

	}

	defer db.Close()

	return user

}

func FindUserByMail(mail string) (user model.User) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		panic(err)

	}

	rows, err := db.Query("SELECT id, mail, username, password FROM user where mail = ? ", mail)

	if err != nil {

		fmt.Println(err)

		os.Exit(1)

	}

	for rows.Next() {

		var id int64
		var username string
		var mail string
		var password string

		err = rows.Scan(&id, &mail, &username, &password)

		user = model.User{
			Id:       id,
			Username: username,
			Mail:     mail,
			Password: password,
		}

	}

	defer db.Close()

	return user

}

func CreateUser(entity model.User) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		panic(err.Error())

	}

	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO user(mail, username, password, token) VALUES(?, ?, ?, ?)")

	if err != nil {

		panic(err.Error())

	}

	uuid, err := exec.Command("uuidgen").Output()

	if err != nil {

		panic(err.Error())

	}

	_, err = stmtIns.Exec(entity.Mail, entity.Username, entity.Password, strings.Replace(string(uuid[:]), "\n", "", -1))

	if err != nil {

		panic(err.Error())

	}

	defer stmtIns.Close()

}

func FindUserByToken(token string) (user model.User) {

	db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")
	//db, err := sql.Open("mysql", "root:Rafilkis1536*@tcp(localhost:3306)/vah_hml?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		panic(err)

	}

	rows, err := db.Query("SELECT id, mail, username, password, token FROM user where token = ? ", token)

	if err != nil {

		fmt.Println(err)

		os.Exit(1)

	}

	for rows.Next() {

		var id int64
		var username string
		var mail string
		var password string
		var token string

		err = rows.Scan(&id, &mail, &username, &password, &token)

		user = model.User{
			Id:       id,
			Username: username,
			Mail:     mail,
			Password: password,
			Token:    token,
		}

	}

	defer db.Close()

	return user

}
