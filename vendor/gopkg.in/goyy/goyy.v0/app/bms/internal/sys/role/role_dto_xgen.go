// generated by xgen -- DO NOT EDIT
package role

import (
	"gopkg.in/goyy/goyy.v0/data/dto"
)

type DTOs []DTO

type DTO struct {
	dto.Sys
	Name    string `json:"name"`
	Code    string `json:"code"`
	Genre   string `json:"genre"`
	Ordinal string `json:"ordinal"`
}