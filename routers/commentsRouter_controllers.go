package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:DatoAdicionalFormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:DatoAdicionalFormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:DatoAdicionalFormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:DatoAdicionalFormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:DatoAdicionalFormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:DatoAdicionalFormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:DatoAdicionalFormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:DatoAdicionalFormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:DatoAdicionalFormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:DatoAdicionalFormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:FormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:FormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:SoporteFormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:SoporteFormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:SoporteFormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:SoporteFormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:SoporteFormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:SoporteFormacionAcademicaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:SoporteFormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:SoporteFormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:SoporteFormacionAcademicaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/formacion_academica_crud/controllers:SoporteFormacionAcademicaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
