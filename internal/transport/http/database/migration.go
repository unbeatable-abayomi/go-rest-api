package database

import (
	"unbeatable-abayomi/go-rest-api/internal/comment"

	"github.com/jinzhu/gorm"
)
//Migerate DB and create comment table
func MigrateDB(db *gorm.DB) error{
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil{
				return result.Error
	}

	return nil
}