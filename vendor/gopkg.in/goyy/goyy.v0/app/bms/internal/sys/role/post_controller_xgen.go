// generated by xgen -- DO NOT EDIT
package role

import (
	"gopkg.in/goyy/goyy.v0/web/controller"
)

var postCtl = &PostController{
	ClientController: controller.ClientController{
		Settings: controller.Settings{
			Project: "sys",
			Module:  "role.post",
			Title:   "ROLE POST",
		},
		DTO: func() interface{} {
			return &PostDTO{}
		},
		DTOs: func() interface{} {
			return &PostDTOs{}
		},
	},
}

type PostController struct {
	controller.ClientController
}
