package service

import "gorm.io/gorm"

type UserService struct {
	db *gorm.DB
}
