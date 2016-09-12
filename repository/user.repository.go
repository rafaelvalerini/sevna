package repository

import( 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"agregador/model"
	"fmt"
	"os"
)

func Login(mail string, password string) (user model.User){

	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err)

	}

	rows, err := db.Query("SELECT id, mail, username FROM user where mail = ? and password = ?", mail, password) 

    if err != nil {

       fmt.Println(err)

        os.Exit(1)

    }

    for rows.Next() {

        var id int64
        var username string
        var mail string
        
        err = rows.Scan(&id, &mail, &username)

        user = model.User{
                Id: id, 
                Username: username, 
                Mail: mail,
            }

	}    
    
    defer db.Close()

    return user

}

func FindUserByMail(mail string) (user model.User){

	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
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
		            Id: id, 
		            Username: username, 
		            Mail: mail,
		            Password: password,
		        }

	}    
    
    defer db.Close()

    return user

}


func CreateUser(entity model.User){

	db, err := sql.Open("mysql", "USR_MOB:mob@money2@tcp(52.87.63.135:3306)/mobint?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {

	    panic(err.Error())

	}

	defer db.Close()
 	
 	stmtIns, err := db.Prepare("INSERT INTO user(mail, username, password) VALUES(?, ?, ?)") 

    if err != nil {

        panic(err.Error())

    }

    _, err = stmtIns.Exec(entity.Mail, entity.Username, entity.Password)

    if err != nil {

    	panic(err.Error())

    }
    
    defer stmtIns.Close()

}

