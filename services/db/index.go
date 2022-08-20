package db

import (
	"fmt"
	"log"

	"github.com/alexmolly/gomvcboilerplate/config"
	"github.com/alexmolly/gomvcboilerplate/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDB(config config.DBConf) *gorm.DB {
	var err error
	conStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(conStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("!! Database Connection Loaded !!")
	return db
}

func DBMigrate(db *gorm.DB) {
	var err error
	for _, model := range model.RegisterModels() {
		err = db.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("!! Database Migration Succeed !!")
}
