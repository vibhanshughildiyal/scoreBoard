package repository

import (
	"gorm.io/gorm"
)

type PlayerDAO interface {
}

type playerDao struct {
	database *gorm.DB
}
