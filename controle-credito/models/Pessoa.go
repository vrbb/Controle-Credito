package models

import (
	"controle-credito/conf"
	"time"

	"github.com/astaxie/beego/orm"
)

type Pessoa struct {
	Cpf            string    `orm:"not_null;unique"`
	Nome           string    `orm:"null"`
	DataNascimento time.Time `orm:"auto_now_add;type(datetime)"`
	DataCriacao    time.Time `orm:"auto_now_add;type(datetime)"`
	conf.Encryptionkey
}

func AddPeople(pessoa Pessoa) *Pessoa {
	o := orm.NewOrm()

	var r orm.RawSeter

	// hash cpf {criptografia}
	// pessoa.Cpf, _ = hashPassword(pessoa.Cpf)

	// get now datetime
	pessoa.DataCriacao = time.Now()
	// Insert
	r = o.Raw("INSERT INTO `Pessoa` (`Cpf`, `Nome`, `DataNascimento`, `DataCriacao`) values(PGP_SYM_ENCRYPT(?,?),PGP_SYM_ENCRYPT(?,?),?,?)",
		pessoa.Cpf,
		pessoa.Encryptionkey,
		pessoa.Nome,
		pessoa.Encryptionkey,
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

func DeletePeople(id int) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&Pessoa{Id: id})
	if err == nil {
		return true
	}

	return false
}

func GetPeople(id int) *Pessoa {
	o := orm.NewOrm()
	pessoa := Pessoa{Id: id}
	o.Read(&pessoa)
	return &pessoa
}
