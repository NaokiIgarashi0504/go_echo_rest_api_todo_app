package main

import (
	"fmt"
	"go_echo_rest_api/models"
	"go_echo_rest_api/routers"
)

func main() {
	fmt.Println(models.Db)

	router := routers.NewMainRouter()
	router.Logger.Fatal(router.Start(":8080"))
}
