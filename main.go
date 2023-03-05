package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kosimaru1997/task_app/config"
)

func main() {

	_, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Ginのルーターを作成する。
	r := gin.Default()

	setRoutes(r)

	// HTTPサーバーを起動する。“
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
