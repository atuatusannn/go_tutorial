package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1/app/models"
)

func main() {
	// fmt.Println("Hello World")
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)
	// u.CreateUser()

	u, _ := models.GetUser(1)
	fmt.Println(u)

	u.Name = "Text2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

}
