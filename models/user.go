package models

import (
	"time" 
)

type User struct {
	ID         *uint   `gorm:"primaryKey" json:"id"`
	Name       *string `gorm:"type:varchar(300)" json:"nama"`
	Email      *string `gorm:"type:varchar(300)" json:"email"`
	Password   *string `gorm:"type:varchar(300)" json:"password"`
	Address    *string `gorm:"type:varchar(300)" json:"address"`
	Level      *uint   ` json:"level"`
	Name_level *string `gorm:"type:varchar(300)" json:"name_level"`
	 CreatedAt time.Time  `gorm:"autoCreateTime"`
	 LastLogin time.Time  `gorm:"default:NULL"`

}
