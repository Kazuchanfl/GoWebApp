package main

import (
	"github.com/Kazuchanfl/GoWebApp/config"
	"github.com/Kazuchanfl/GoWebApp/config/router"
)

func main() {

	// DBの起動
	config.SetupDb()

	// gin web server を起動
	router.SetupRouter()

}
