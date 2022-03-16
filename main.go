package main

import (
	"fmt"
	"memo_app/app/controllers"
	"memo_app/app/models"
)

func main() {

	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// log.Println("test")

	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)
	// u.CreateUser()

	// user, _ := models.GetUser(1)
	// fmt.Println(user)

	// user.CreateTodo("First Todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user, _ := models.GetUser(1)
	// todos, _ := user.GetTodoByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	controllers.StartMainServer()

}
