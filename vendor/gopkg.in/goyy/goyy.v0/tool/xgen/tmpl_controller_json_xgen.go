// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplControllerJsonXgen = `// generated by xgen -- DO NOT EDIT
package {{.PackageName}}{{with $ := .}}

import (
	"gopkg.in/goyy/goyy.v0/web/controller"
){{range $i, $e := .Entities}}{{with $name := ctlname $e.Name $e.Relationship}}

var {{ctlvar $e.Name $e.Relationship}} = &{{$name}}{
	{{if eq "tree" $e.Extend}}JSONTreeController: controller.JSONTreeController{
		JSONController: controller.JSONController{
			Settings: controller.Settings{
				Project: "{{$e.Project}}",
				Module:  "{{mname $e.Name $.PackageName}}",
				Title:   "{{$e.Comment}}",
			},
			Mgr: {{mgrvar $e.Name $e.Relationship}},
		},
	},{{else}}JSONController: controller.JSONController{
		Settings: controller.Settings{
			Project: "{{$e.Project}}",
			Module:  "{{mname $e.Name $.PackageName}}",
			Title:   "{{$e.Comment}}",
		},
		Mgr: {{mgrvar $e.Name $e.Relationship}},
	},{{end}}
}

type {{$name}} struct {
	{{if eq "tree" $e.Extend}}controller.JSONTreeController{{else}}controller.JSONController{{end}}
}{{end}}{{end}}{{end}}
`
