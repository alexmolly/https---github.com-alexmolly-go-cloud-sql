package services

import (
	"github.com/alexmolly/gomvcboilerplate/services/example"
	"gorm.io/gorm"
)

var (
	ExampleService example.ExampleService
)

func InjectDBIntoServices(db *gorm.DB) {
	ExampleService.DB = db
}
