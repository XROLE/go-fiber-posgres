package models

import "gorm.io/gorm"

type Books struct {
	ID 				uint 			`gorm:"primary key;autoIncrement" json:"id"`
	Authur 			*string			`json:"author"`
	Title 			*string 		`json:"title"`
	Publisher		*string 		`json:"publisher"`

}

func MigrateBooks (db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})

	return err
}