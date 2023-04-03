package models

type File struct {
	Id       uint   `gorm:"autoIncrement;primary_key"`
	Url      string `gorm:"type:string"`
	FileType string `gorm:"type:varchar(50)"`
}
