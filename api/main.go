package main

import (
	"fiber-stats/bookroute"
	"fiber-stats/util"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	data := util.Data{}
	_, err := data.LoadConfiguration("util/config.json")
	if err != nil {
		log.Fatalln("Failed at config", err.Error())
	}

	port := data.GetInt("port")

	dbUser := data.GetString("database.user")
	dbPass := data.GetString("database.password")
	dbHost := data.GetString("database.host")
	dbPort := data.GetString("database.port")
	dbName := data.GetString("database.dbname")

	connectionStringURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	app := fiber.New()
	db := util.Init(connectionStringURL)

	bookroute.RegisterRoutes(app, db)

	app.Listen(fmt.Sprintf(":%d", port))

}
