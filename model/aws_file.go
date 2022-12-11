package model

import "time"

type Aws_File struct {
	FileId    string    `json:"file_id" gorm:"primaryKey" gorm:"column:file_id" validate:"required"`
	AwsLink   string    `json:"aws_link" gorm:"column:aws_link"`
	Extension string    `json:"extension" gorm:"column:extension"`
	Password  string    `json:"password" gorm:"column:password"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}
