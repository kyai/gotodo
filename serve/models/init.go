package models

import (
	"gotodo/serve/core"
	"gotodo/serve/flags"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {
	log.Println("SQLite connecting...")

	db, err := gorm.Open(sqlite.Open(flags.Data+"/todo.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	core.Bomb(err)

	sqldb, err := db.DB()
	core.Bomb(err)
	core.Bomb(sqldb.Ping())

	if flags.Debug {
		DB = db.Debug()
	} else {
		DB = db
	}
}
