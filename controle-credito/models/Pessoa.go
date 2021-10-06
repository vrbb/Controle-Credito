package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Pessoa struct {
	Id             int
	Cpf            string    `orm:"null;unique"`
	Nome           string    `orm:"null"`
	DataNascimento time.Time `orm:"auto_now_add;type(datetime)"`
	DataCriacao    time.Time `orm:"auto_now_add;type(datetime)"`
}

func main() {
	orm.RegisterModel(new(Pessoa))

	o := orm.NewOrm()

	pessoa := new(Pessoa)

	fmt.Println(o.Insert(pessoa))

	pessoa.Nome = "Your"
	fmt.Println(o.Update(pessoa))
	fmt.Println(o.Read(pessoa))
	fmt.Println(o.Delete(pessoa))

	// insert
	id, err := o.Insert(&pessoa)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	pessoa.Nome = "astaxie"
	num, err := o.Update(&pessoa)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := Pessoa{Id: pessoa.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}

func InserirPessoa(pessoa Pessoa) *Pessoa {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Pessoa))

	i, _ := qs.PrepareInsert()

	var p Pessoa

	// hash cpf {criptografia}
	// pessoa.Cpf, _ = hashPassword(pessoa.Cpf)

	// get now datetime
	pessoa.DataCriacao = time.Now()

	// Insert
	id, err := i.Insert(&pessoa)
	if err == nil {
		// successfully inserted
		p = Pessoa{Id: int(id)}
		err := o.Read(&p)
		if err == orm.ErrNoRows {
			return nil
		}
	} else {
		return nil
	}

	return &p
}

func BuscarTodasPessoas() []*Pessoa {
	o := orm.NewOrm()
	var pessoas []*Pessoa
	o.QueryTable(new(Pessoa)).All(&pessoas)

	return pessoas
}

func AtualizarPessoa(pessoa Pessoa) *Pessoa {
	o := orm.NewOrm()
	p := Pessoa{Id: pessoa.Id}
	var atualizardPessoa Pessoa

	// get existing pessoa
	if o.Read(&p) == nil {

		pessoa.DataCriacao = p.DataCriacao
		p = pessoa
		_, err := o.Update(&p)

		// read updated pessoa
		if err == nil {
			// update successful
			atualizardPessoa = Pessoa{Id: pessoa.Id}
			o.Read(&atualizardPessoa)
		}
	}

	return &atualizardPessoa
}

func DeletarPessoa(id int) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&Pessoa{Id: id})
	if err == nil {
		return true
	}

	return false
}

func BuscarPessoaPeloId(id int) *Pessoa {
	o := orm.NewOrm()
	pessoa := Pessoa{Id: id}
	o.Read(&pessoa)
	return &pessoa
}
