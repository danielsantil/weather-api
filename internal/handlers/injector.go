package handlers

import "gorm.io/gorm"

type Injector struct {
	DB *gorm.DB
}
