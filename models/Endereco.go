package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Endereco struct {
	Id              int       `orm:"not_null;unique"`
	Logadouro       string    `orm:"null"`
	Numero          string    `orm:"null"`
	Bairro          string    `orm:"null"`
	Cidade          string    `orm:"null"`
	Estado          string    `orm:"null"`
	Complemento     string    `orm:"null"`
	PontoReferencia string    `orm:"null"`
	DataCriacao     time.Time `orm:"auto_now_add;type(datetime)"`
	Cpf             string    `orm:"not_null;unique"`
}

func AddAddress(endereco Endereco) *Endereco {
	o := orm.NewOrm()

	var r orm.RawSeter

	// hash cpf {criptografia}
	// endereco.Cpf, _ = hashPassword(endereco.Cpf)

	// get now datetime
	endereco.DataCriacao = time.Now()
	// Insert
	r = o.Raw("INSERT INTO `Endereco` (`Logadouro`, `Numero`, `Bairro`, `Cidade`, `Estado`, `Estado`, `Complemento`, `PontoReferencia`, `DataCriacao`) values(?,?,?,?,?,?,?,?)",
		endereco.Logadouro,
		endereco.Numero,
		endereco.Bairro,
		endereco.Cidade,
		endereco.Estado,
		endereco.Complemento,
		endereco.PontoReferencia,
		endereco.DataCriacao)
	_, err := r.Exec()

	if err == nil {

		err := o.Read(&endereco)
		if err == orm.ErrNoRows {
			return nil
		}
	} else {
		return nil
	}

	return &endereco
}

func AllAddress() []*Endereco {
	o := orm.NewOrm()
	var enderecos []*Endereco
	o.QueryTable(new(Endereco)).All(&enderecos)

	return enderecos
}

func UpdateAddress(endereco Endereco) *Endereco {
	o := orm.NewOrm()
	e := Endereco{Id: endereco.Id}
	var atualizardEndereco Endereco

	// get existing endereco
	if o.Read(&e) == nil {

		endereco.DataCriacao = e.DataCriacao
		e = endereco
		_, err := o.Update(&e)

		// read updated endereco
		if err == nil {
			// update successful
			atualizardEndereco = Endereco{Id: endereco.Id}
			o.Read(&atualizardEndereco)
		}
	}

	return &atualizardEndereco
}

func DeleteAddress(id int) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&Endereco{Id: id})
	if err == nil {
		return true
	}

	return false
}

func GetAddress(id int) *Endereco {
	o := orm.NewOrm()
	endereco := Endereco{Id: id}
	o.Read(&endereco)
	return &endereco
}
