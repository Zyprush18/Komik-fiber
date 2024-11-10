package entity

import "time"


type Komik struct{
	Id uint `json:"id" grom:"primaryKey;"`
	Name string `json:"name_comic" grom:"type:varchar(255)"`
	Creator string `json:"name_creator" grom:"type:varchar(255)"`
	Publisher string `json:"name_publisher" grom:"type:varchar(255)"`
	Releases string `json:"release_comic" grom:"type:varchar(255)"`
	Image string `json:"image_comic" gorm:"text"`
	Decsription string `json:"description" gorm:"text"`
	CreatedAt time.Time
	UpdatedAt time.Time

}