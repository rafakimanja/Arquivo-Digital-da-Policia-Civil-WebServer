package models

import "gorm.io/gorm"

type Sistema struct {
	gorm.Model
	TipoArq     string `json:"tipo_arq"`
	UploadUnico bool   `json:"upload_unico"`
}

var ConfSistema Sistema