package model

type Tags struct {
	ID   int    `gorm:"primary_key;AUTO_INCREMENT;column:ID"`
	Name string `gorm:"column:Name"`
}
