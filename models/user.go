package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username	string	`json:"username" gorm:"unique;not null"`
	Email    	string	`json:"email" gorm:"unique;not null"`
	FirstName	string	`json:"first_name" gorm:"type:varchar(255)"`
	LastName	string	`json:"last_name" gorm:"type:varchar(255)"`
	Password 	string	`json:"password" gorm:"not null"`
	PerPhone	string	`json:"per_phone" gorm:"type:varchar(255)"`
	AltPhone	string	`json:"alt_phone" gorm:"type:varchar(255)"`
	Balance		string	`json:"balance" gorm:"type:decimal(12,6)"`
	Address		string	`json:"address" gorm:"type:varchar(255)"`
	City		string	`json:"city" gorm:"type:varchar(255)"`
	Country		string	`json:"country" gorm:"type:varchar(255)"`
	Pincode		string	`json:"pincode" gorm:"type:varchar(255)"`
}
