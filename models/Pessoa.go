package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

const Encryptionkey = "d41d8cd98f00b204e9800998ecf8427e"

type Pessoa struct {
	Cpf            string    `orm:";unique"`
	Nome           string    `orm:"null"`
	DataNascimento time.Time `orm:";type(datetime)"`
	DataCriacao    time.Time `orm:"auto_now_add;type(datetime)"`
}

func (P *Pessoa) AddPeople(pessoa Pessoa) *Pessoa {
	o := orm.NewOrm()
	fmt.Println("Entrou patroa")
	var r orm.RawSeter

	// hash cpf {criptografia}
	// pessoa.Cpf, _ = hashPassword(pessoa.Cpf)

	// get now datetime
	pessoa.DataCriacao = time.Now()
	// Insert
	r = o.Raw("INSERT INTO `pessoa` (`Cpf`, `Nome`, `DataNascimento`, `DataCriacao`) values(PGP_SYM_ENCRYPT(?,?),PGP_SYM_ENCRYPT(?,?),?,?)",
		pessoa.Cpf,
		Encryptionkey,
		pessoa.Nome,
		Encryptionkey,
		pessoa.DataNascimento,
		pessoa.DataCriacao)
	_, err := r.Exec()

	if err == nil {

		err := o.Read(&pessoa)
		if err == orm.ErrNoRows {
			return nil
		}
	} else {
		return nil
	}

	return &pessoa
}

func AllPeople() []*Pessoa {
	o := orm.NewOrm()
	var pessoas []*Pessoa
	o.QueryTable(new(Pessoa)).All(&pessoas)

	return pessoas
}

func UpdatePeople(pessoa Pessoa) *Pessoa {
	o := orm.NewOrm()
	p := Pessoa{Cpf: pessoa.Cpf}
	var atualizardPessoa Pessoa

	// get existing pessoa
	if o.Read(&p) == nil {

		pessoa.DataCriacao = p.DataCriacao
		p = pessoa
		_, err := o.Update(&p)

		// read updated pessoa
		if err == nil {
			// update successful
			atualizardPessoa = Pessoa{Cpf: pessoa.Cpf}
			o.Read(&atualizardPessoa)
		}
	}

	return &atualizardPessoa
}

func DeletePeople(cpf string) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&Pessoa{Cpf: cpf})
	if err == nil {
		return true
	}

	return false
}

func GetPeople(cpf string) *Pessoa {
	o := orm.NewOrm()
	pessoa := Pessoa{Cpf: cpf}
	o.Read(&pessoa)
	return &pessoa
}
