package main_test

import (
	"log"
	"os"
	"testing"

	"github.com/eRaBxEs/go-fiber-api/pkg/books"
	"github.com/eRaBxEs/go-fiber-api/pkg/common/config"
	"github.com/eRaBxEs/go-fiber-api/pkg/common/db"
	"github.com/eRaBxEs/go-fiber-api/pkg/common/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := db.Init(c.DBUrl)
	app := fiber.New()

	ensureTableExists(db)
	code := m.Run()
	util.ClearTable(db)
	books.RegisterRoutes(app, db)

	os.Exit(code)
}

func ensureTableExists(db *gorm.DB) {
	if err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS books
(
    id SERIAL,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
	description TEXT NOT NULL,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`
