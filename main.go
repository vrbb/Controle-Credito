package main

import (
	"controle-credito/models"
	_ "controle-credito/routers"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// func init() {
// 	// // registra os modelos

// 	orm.RegisterModel(new(models.Endereco))
// 	orm.RegisterModel(new(models.BemMaterial))
// 	orm.RegisterModel(new(models.Divida))
// 	orm.RegisterModel(new(models.Renda))

// 	// // registra o drive
// 	orm.RegisterDriver("postgres", orm.DRPostgres)

// 	// // configura o pool de conexões {Usar Enable quando for usar a criptografia}
// 	orm.RegisterDataBase("default",
// 		"postgres",
// 		"postgres://postgres:123456example@172.17.0.1:5432/postgres?sslmode=disable")

// 	// cria as tabelas
// 	if err := orm.RunSyncdb("default", false, true); err != nil {
// 		fmt.Println(err)
// 	}
// 	// Pessoa := new(models.Pessoa)
// 	// P := new(models.Pessoa)
// 	// P.Cpf = "06262874305"
// 	// P.Nome = "Victoria Ricarte Bispo Beserra"
// 	// P.DataCriacao = time.Now()
// 	// Pessoa.AddPeople(*P)

// }

func init() {
	// registra os modelos
	orm.RegisterModel(new(models.Pessoa))
	orm.RegisterModel(new(models.Endereco))
	orm.RegisterModel(new(models.BemMaterial))
	orm.RegisterModel(new(models.Divida))
	orm.RegisterModel(new(models.Renda))

	// registra o drive
	orm.RegisterDriver("postgres", orm.DRPostgres)

	// configura o pool de conexões {Usar Enable quando for usar a criptografia}
	orm.RegisterDataBase("default",
		"postgres",
		"user=postgres password=root host=127.0.0.1 port=5432 dbname=credito sslmode=disable")

	// cria as tabelas
	if err := orm.RunSyncdb("default", false, true); err != nil {
		fmt.Println(err)
	}
}

func main() {
	beego.Run()

}
