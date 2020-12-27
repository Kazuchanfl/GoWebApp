package main

import (
	"github.com/Kazuchanfl/GoWebApp/config"
)

func main() {

	// DBの起動
	config.SetupDb()

	// gin web server を起動
	config.SetupRouter()

}
