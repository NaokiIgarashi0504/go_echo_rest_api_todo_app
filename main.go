package main

import (
	"fmt"
	"go_echo_rest_api/db"
)

func main() {
	// DB接続
	db := db.Init()

	fmt.Println(db)

	// 下記3行はDIを踏まえた実装前のコード
	// fmt.Println(models.Db)
	// router := routers.NewMainRouter()
	// router.Logger.Fatal(router.Start(":8080"))
}
