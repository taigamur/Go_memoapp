package main

import (
	"fmt"
	"udemy1_todo_app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// log.Println("test")

	// fmt.Println(models.Db)

	/*
		u := &models.User{}
		u.Name = "test"
		u.Email = "test@example.com"
		u.PassWord = "testtest"
		fmt.Println(u)
		u.CreateUser()
	*/

	u, err := models.GetUser(1)
	fmt.Println(u)
	fmt.Println(err)

	u.Name = "test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}
