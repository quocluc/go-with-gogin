package main

import (
	"apiProject/go-gin/models"
	"apiProject/go-gin/routers"
)

func main() {

	models.Setup()
	router := routers.InitRouter()
	router.Run()
	models.CloseDB()
}
