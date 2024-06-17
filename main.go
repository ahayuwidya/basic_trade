package main

import (
	"basic_trade/database"
	"basic_trade/routers"
)

func main() {
	database.InitDB()
	database.GetDB()

	var PORT = ":8080"
	routers.StartApp().Run(PORT)
}
