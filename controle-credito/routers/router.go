package routers

import (
	"controle-credito/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// Pessoa

	// Endereço

	// Bem material

	// Dívida

	// Renda
}
