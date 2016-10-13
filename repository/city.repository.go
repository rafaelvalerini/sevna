package repository

import( 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"agregador/model"
	"fmt"
)

func FindAllStates() (states []model.State){

    db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
    
    if err != nil {

        panic(err)

    }

    rows, err := db.Query("SELECT id, nome, uf FROM estado") 

    if err != nil {

    	defer db.Close()

        fmt.Println(err)

        return states;

    }

    for rows.Next() {

        var id int

        var name string

        var uf string
        
        err = rows.Scan(&id, &name, &uf)

        states = append(states,  model.State{Id: id, Name: name, Initials: uf,})

    }    
    
    defer db.Close()

    return states

}

func FindCityByState(state int) (cities []model.City){

    db, err := sql.Open("mysql", "usr_vah:vah_taxi2$@tcp(vah.cn73hi7irhmm.us-east-1.rds.amazonaws.com:3306)/vah?charset=utf8&parseTime=True&loc=Local")
    
    if err != nil {

        panic(err)

    }

    rows, err := db.Query("SELECT id, nome FROM cidade where estado = ?", state) 

    if err != nil {

    	defer db.Close()

        fmt.Println(err)

        return cities;

    }

    for rows.Next() {

        var id int

        var name string
        
        err = rows.Scan(&id, &name)

        cities = append(cities,  model.City{Id: id, Name: name,})

    }    
    
    defer db.Close()

    return cities

}



