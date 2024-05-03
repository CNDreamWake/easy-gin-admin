package model

import "gorm.io/gorm"

// InitializeDatabaseSchema 迁移数据表结构
func InitializeDatabaseSchema(db *gorm.DB) error {
	return db.AutoMigrate(
		&T{})
}
