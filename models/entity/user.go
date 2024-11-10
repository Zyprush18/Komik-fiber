package entity

import (
	"time"
)

type User struct{
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" grom:"type:varchar(255)"`
	Email string `json:"email" grom:"type:varchar(255)"`
	Password string `json:"password" grom:"type:varchar(255)"`
	CreatedAt time.Time
  	UpdatedAt time.Time
  	// DeletedAt gorm.DeletedAt `json:"-"  gorm:"index, column:deleted_at"`
}