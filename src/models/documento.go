package models

import "gorm.io/gorm"

type Documento struct {
	gorm.Model
	Nome      string 
	Ano       int    
	Categoria string 
	Arquivo   string 
}
