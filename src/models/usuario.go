package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome  string `json:"nome"`
	RG    string `json:"rg"`
	Senha string `json:"senha"`
	Admin bool   `json:"admin"`
}
