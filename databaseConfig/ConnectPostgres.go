package databaseConfig

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var Config = viper.New()

func ConnectToDb() *sqlx.DB {
	Config.SetConfigFile("config.yaml")
	Config.ReadInConfig()
	db, err := ConnectDB(
		Config.GetString("db.host"),
		Config.GetString("db.port"),
		Config.GetString("db.dbname"),
		Config.GetString("db.username"),
		Config.GetString("db.password"),
		Config.GetString("db.sslmode"),
	)
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	log.Println("Database connection success!")

	return db
}

func ConnectDB(host, port, dbname, user, password, sslmode string) (*sqlx.DB, error) {
	databaseURL := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host,
		port,
		dbname,
		user,
		password,
		sslmode,
	)
	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
