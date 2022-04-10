package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)	

func main() {

	//configの設定
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	//user作成
	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)

	// fmt.Println(u)


	//user編集
	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()

	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	//user削除
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)




	//todo作成
	// user, _ := models.GetUser(2)
	// user.CreateTodo("First Todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third Todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	//todo取得(1ユーザー)
	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetTodosByUser()
	// 	for _, v := range todos {
	// 	fmt.Println(v)
	// }


	//todo編集
	// t, _ := models.GetTodo(1)
	// t.Content = "Update Todo"
	// t.UpdateTodo()

	//todo削除
	// t, _ := models.GetTodo(3)
	// t.DeleteTodo()

	//サーバーの起動
	controllers.StartMainServer()


	//loginセッション使用
	// user, _ := models.GetUserByEmail("test@exmaple.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
}