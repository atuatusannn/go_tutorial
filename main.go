package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1/config"
	// "gopkg.in/go-ini/ini.v1/config"
	// "gopkg.in/go-ini/ini.v1/config"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.LogFile)

}
