package router

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TestRegisterRoutes(t *testing.T) {
	type args struct {
		app *fiber.App
		db  *gorm.DB
	}

	argsTest := args{nil, nil}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "failed routes", args: argsTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterRoutes(tt.args.app, tt.args.db)
		})
	}
}
