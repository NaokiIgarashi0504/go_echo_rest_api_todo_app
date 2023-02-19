package main

import (
	"go_echo_rest_api/controllers"
	"go_echo_rest_api/db"
	"go_echo_rest_api/repositories"
	"go_echo_rest_api/routers"
	"go_echo_rest_api/services"
)

func main() {
	// DB接続
	db := db.Init()

	// repository層
	authRepo := repositories.NewAuthRepository(db)
	todoRepo := repositories.NewTodoRepository(db)
	// userRepo := repositories.NewAuthRepository(db)
	// todoRepo := repositories.NewTodoRepository(db)

	// service層
	authService := services.NewAuthService(authRepo)
	todoService := services.NewTodoService(todoRepo, authRepo)
	baseService := services.NewBaseService(authRepo)
	// authService := services.NewAuthService(authRepo)
	// todoService := services.NewTodoService(todoRepo, todoLogic, responseLogic, todoValidate)

	// controller層
	authController := controllers.NewAuthController(authService, baseService)
	todoController := controllers.NewTodoController(todoService, authService, baseService)
	// appController := controllers.NewAppController()
	// authController := controllers.NewAuthController(authService)
	// todoContoroller := controllers.NewTodoController(todoService, authService)

	// router設定
	authRouter := routers.NewAuthRouter(authController)
	todoRouter := routers.NewTodoRouter(todoController)
	mainRouter := routers.NewMainRouter(authRouter, todoRouter)
	// appRouter := router.NewAppRouter(appController)
	// authRouter := router.NewAuthRouter(authController)
	// todoRouter := router.NewTodoRouter(todoContoroller)
	// mainRouter := router.NewMainRouter(appRouter, authRouter, todoRouter)

	// 起動
	mainRouter.StartWebServer()

	// router := mainRouter.StartWebServer()
	// fmt.Println(router)
	// router.Start(":8080")
	// 音量呂布カルマどっち通.StartWebServer()音量呂布カルマどっち

	// 下記3行はDIを踏まえた実装前のコード
	// fmt.Println(models.Db)
	// router := routers.NewMainRouter()
	// router.Logger.Fatal(router.Start(":8080"))
}
