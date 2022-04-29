package router

import (
	"fiber-stats/util"
	"fmt"
	"log"
	"testing"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Test_handler_SlowQuery(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		c *fiber.Ctx
	}

	data := util.Data{}
	_, err := data.LoadConfiguration("./util/config.json")
	if err != nil {
		log.Fatalln("Failed at config", err.Error())
	}

	config := util.ConnConfig{
		DBUser: data.GetString("database.user"),
		DBPass: data.GetString("database.password"),
		DBHost: data.GetString("database.host"),
		DBPort: data.GetString("database.port"),
		DBName: data.GetString("database.dbname"),
	}

	connectionStringURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)

	// a := fiber.New()
	// app := a.AcquireCtx(nil)
	db := util.Init(connectionStringURL)

	correctField := fields{DB: db}
	correctArgs := args{c: nil}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Successful handler", fields: correctField, args: correctArgs, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				DB: tt.fields.DB,
			}
			if err := h.SlowQuery(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("handler.SlowQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
