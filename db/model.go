package db

import "gorm.io/gorm"

type Session struct {
	gorm.Model

	OpenId         string `gorm:"column:openId;index"`
	MessageHistory []byte `gorm:"column:messageHistory"`
}
