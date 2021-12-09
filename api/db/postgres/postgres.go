package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dgkg/gomasterclasses/api/db/sqlite"
	"github.com/dgkg/gomasterclasses/api/model"
)

type PostgresDB = sqlite.SQLite

func New() *PostgresDB {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = conn.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	var pg PostgresDB
	pg.SetConn(conn)
	return &pg
}
