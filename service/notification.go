package service

import (
	"agregador/model"
	"github.com/alexjlockwood/gcm"
	"agregador/repository"
	"fmt"
	"os"
)

const (

	KEY_NOTIFICATION = "AIzaSyCOR9rwLoyX6WZDkX-gTQJM-ozTi63x8qM"

)

func SendNotification(state string, city string, message string) (meta model.Meta) {

 	go func() {

		data := map[string]interface{}{"message": message}

		page := 0

		notifications := repository.SearchByNotification(state, city, page)

		for len(notifications) > 0 {

			msg := gcm.NewMessage(data, notifications...)

			sender := &gcm.Sender{ApiKey: KEY_NOTIFICATION}

			_, err := sender.Send(msg, 2)

			if err != nil{

				fmt.Println(err)
 				
 				os.Exit(1)

			}

			page = page + 1

			notifications = repository.SearchByNotification(state, city, (page * 100))

		} 

	}()

	return model.Meta{Value: "OK"}

}