package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

type Database struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewDatabase(logger *slog.Logger) *Database {
	dsn := viper.GetString("database.dsn")
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(viper.GetInt("database.max_idle_conn"))
	db.SetMaxOpenConns(viper.GetInt("database.max_open_conn"))
	if err = db.Ping(); err != nil {
		logger.Error("ping error", "Connection error", err)
		panic(err)
	}
	return &Database{
		logger: logger,
		db:     db,
	}
}
