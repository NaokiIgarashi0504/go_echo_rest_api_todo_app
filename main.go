package main

import (
	"fmt"
	"go_echo_rest_api/db"
	"go_echo_rest_api/repositories"
	"go_echo_rest_api/services"
)

func main() {
	// DB接続
	db := db.Init()

	// repository層
	authRepo := repositories.NewAuthRepository(db)
	// userRepo := repositories.NewAuthRepository(db)
	// todoRepo := repositories.NewTodoRepository(db)

	// service層
	authService := services.NewAuthService(authRepo)
	// authService := services.NewAuthService(authRepo)
	// todoService := services.NewTodoService(todoRepo, todoLogic, responseLogic, todoValidate)

	fmt.Println(authRepo)
	fmt.Println(authService)

	// 下記3行はDIを踏まえた実装前のコード
	// fmt.Println(models.Db)
	// router := routers.NewMainRouter()
	// router.Logger.Fatal(router.Start(":8080"))
}
