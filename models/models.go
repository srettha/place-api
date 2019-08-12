package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// For database connection
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/thestrayed/place-api/config"
)

// db :: Database connection
var db *gorm.DB

// Init :: Initialize connection with Database
func Init() error {
	dbURI := fmt.Sprintf(
		"sslmode=disable host=%s user=%s dbname=%s password=%s",
		config.Values.Database.Host,
		config.Values.Database.Username,
		config.Values.Database.Name,
		config.Values.Database.Password,
	)

	conn, err := gorm.Open(config.Values.Database.Dialect, dbURI)
	if err != nil {
		return err
	}

	db = conn

	db.AutoMigrate(&Place{})

	return nil
}
