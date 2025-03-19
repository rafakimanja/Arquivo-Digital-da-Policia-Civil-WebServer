package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome     	string    `json:"nome"`
	Usuario       string    `json:"usuario"`
	Senha    string    `json:"senha"`
	Admin    bool      `json:"admin"`
	Sessions []Session `gorm:"foreignKey:UserID"`
}
