package handlers

import "gorm.io/gorm"

// Injector holds references that can be used for all handlers.
type Injector struct {
	DB *gorm.DB
}
