package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct{
	Id uint `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time
  	UpdatedAt time.Time
  	DeletedAt gorm.DeletedAt `gorm:"index"`
}