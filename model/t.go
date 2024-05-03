package model

type T struct {
	Model
	Id uint `gorm:"primaryKey"`
}

func (T) TableName() string {
	return "t"
}
