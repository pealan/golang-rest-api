package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDSN() string {
	// TODO read TEST/PROD config file and appropriately build DSN
	return "host=golang-rest-api-db user=pg password=pass dbname=golang-rest-api-db port=5432"
}

func ConnectToDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getDSN()), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database: %+v", err)
		return nil, err
	}

	return db, nil
}
