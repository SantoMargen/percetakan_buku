package main

import (
	"siap_app/internal/app"

	"github.com/sirupsen/logrus"
)

func main() {
	application := app.NewApp()
	defer application.DB.Close()
	defer application.Redis.Close()

	err := application.Run(":8080")
	if err != nil {
		logrus.Fatal("Failed to start the server: ", err)
		panic(err)
	}
}
