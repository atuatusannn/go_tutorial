package main

import (
	"fmt"
	"log"

	"gopkg.in/go-ini/ini.v1/app/models"
)

func main() {
	fmt.Println(models.Db)

	// controllers.StartMainServer()
	user, _ := models.GetUserByEmail("test@mail.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)
}
