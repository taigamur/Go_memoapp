package main

import (
	"fmt"
	"udemy1_todo_app/app/models"
	"udemy1_todo_app/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	// log.Println("test")

	fmt.Println(models.Db)

}
