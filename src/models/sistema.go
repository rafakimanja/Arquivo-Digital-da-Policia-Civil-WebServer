package models

import "gorm.io/gorm"

type Sistema struct {
	gorm.Model
	TipoArq     string 
	UploadUnico bool   
}

var ConfSistema Sistema