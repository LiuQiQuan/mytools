package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["mytools/controllers:HostRegistryController"] = append(beego.GlobalControllerRouter["mytools/controllers:HostRegistryController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["mytools/controllers:HostRegistryController"] = append(beego.GlobalControllerRouter["mytools/controllers:HostRegistryController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["mytools/controllers:HostRegistryController"] = append(beego.GlobalControllerRouter["mytools/controllers:HostRegistryController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["mytools/controllers:HostRegistryController"] = append(beego.GlobalControllerRouter["mytools/controllers:HostRegistryController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["mytools/controllers:HostRegistryController"] = append(beego.GlobalControllerRouter["mytools/controllers:HostRegistryController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["mytools/controllers:RdpDestopController"] = append(beego.GlobalControllerRouter["mytools/controllers:RdpDestopController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["mytools/controllers:RdpDestopController"] = append(beego.GlobalControllerRouter["mytools/controllers:RdpDestopController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["mytools/controllers:RdpDestopController"] = append(beego.GlobalControllerRouter["mytools/controllers:RdpDestopController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["mytools/controllers:RdpDestopController"] = append(beego.GlobalControllerRouter["mytools/controllers:RdpDestopController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["mytools/controllers:RdpDestopController"] = append(beego.GlobalControllerRouter["mytools/controllers:RdpDestopController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
