package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username	string	`gorm:"unique;not null"`
	Email    	string	`gorm:"unique;not null"`
	FirstName	string	`gorm:"type:varchar(255)"`
	LastName	string	`gorm:"type:varchar(255)"`
	Password 	string	`gorm:"not null"`
}
