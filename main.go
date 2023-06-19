package main

import (
	"pos-echo/db"
	"pos-echo/router"
)

func main() {
	db.NewDB()
	e := router.Init()
	e.Logger.Fatal(e.Start(":1323"))

}
