package util

import (
	"gorm.io/gorm"
)

func ClearTable(db *gorm.DB) {
	db.Exec("DELETE FROM books")
	db.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")
}
