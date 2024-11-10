package entity

import (
	"time"
)

type User struct{
	Id uint `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time
  	UpdatedAt time.Time
  	// DeletedAt gorm.DeletedAt `json:"-"  gorm:"index, column:deleted_at"`
}