package handler

import (
	"gorm.io/gorm"
)

type HandlerDeps struct {
	DB *gorm.DB
}
