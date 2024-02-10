package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `gorm:"not null;unique"`
	Email      string `gorm:"not null;unique"`
	Password   string `gorm:"not null"`
	ActiveCode uint   `sql:"DEFAULT:0"`
	IsActive   bool   `sql:"DEFAULT:false"`
	ImageName  string `sql:"DEFAULT:default_avatar.jpg"`
}
