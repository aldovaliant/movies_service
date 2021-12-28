package config

import (
	"fmt"
	"log"
	"movies/entities/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	dbHost := ENV.DbHost
	dbPort := ENV.DbPort
	dbUser := ENV.DbUser
	dbPass := ENV.DbPass
	dbName := ENV.DbName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	autoCreate := ENV.DbMigrate
	if autoCreate == "true" {
		fmt.Println("Dropping and recreating all tables")
		dbConn.AutoMigrate(
			&schema.Logs{},
		)
		fmt.Println("All tables recreated successfully")
	}

	return dbConn
}
