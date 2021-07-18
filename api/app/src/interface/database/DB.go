package database

import "github.com/jinzhu/gorm"

type DB struct {
	Connect *gorm.DB
}
