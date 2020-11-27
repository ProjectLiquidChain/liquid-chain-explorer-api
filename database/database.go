package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database retains the connection to database server
type Database struct {
	*gorm.DB
}

// New returns new instance of database
func New(url string) Database {
	postgresDB, err := gorm.Open(postgres.Open(url), &gorm.Config{Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second * 5,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)})
	if err != nil {
		log.Fatal(err)
	}
	db := Database{postgresDB}
	db.migrate()
	return db
}
