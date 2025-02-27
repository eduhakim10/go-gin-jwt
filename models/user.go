package models

type User struct {
	ID         *uint   `gorm:"primaryKey" json:"id"`
	Nama       *string `gorm:"type:varchar(300)" json:"nama"`
	Email      *string `gorm:"type:varchar(300)" json:"email"`
	Password   *string `gorm:"type:varchar(300)" json:"password"`
	Address    *string `gorm:"type:varchar(300)" json:"address"`
	Level      *uint   ` json:"level"`
	Nama_level *string `gorm:"type:varchar(300)" json:"nama_level"`
}
