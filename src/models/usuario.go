package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome     string    
	RG       string    
	Senha    string    
	Admin    bool      
	Sessions []Session `gorm:"foreignKey:UserID"`
}
