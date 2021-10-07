package models

import (
	"controle-credito/conf"
	"time"

	"github.com/astaxie/beego/orm"
)

type Divida struct {
	Id          int       `orm:"not_null;unique"`
	Valor       float32   `orm:"null"`
	DataCriacao time.Time `orm:"auto_now_add;type(datetime)"`
	Cpf         string    `orm:"not_null;unique"`
	conf.Encryptionkey
}

func AddDebt(divida Divida) *Divida {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Divida))

	i, _ := qs.PrepareInsert()

	var d Divida

	// hash cpf {criptografia}
	// divida.Cpf, _ = hashPassword(divida.Cpf)

	// get now datetime
	divida.DataCriacao = time.Now()

	// Insert
	id, err := i.Insert(&divida)
	if err == nil {
		// successfully inserted
		d = Divida{Id: int(id)}
		err := o.Read(&d)
		if err == orm.ErrNoRows {
			return nil
		}
	} else {
		return nil
	}

	return &d
}

func AllDebt() []*Divida {
	o := orm.NewOrm()
	var dividas []*Divida
	o.QueryTable(new(Divida)).All(&dividas)

	return dividas
}

func UpdateDebt(divida Divida) *Divida {
	o := orm.NewOrm()
	d := Divida{Id: divida.Id}
	var atualizardDivida Divida

	// get existing divida
	if o.Read(&d) == nil {

		divida.DataCriacao = d.DataCriacao
		d = divida
		_, err := o.Update(&d)

		// read updated divida
		if err == nil {
			// update successful
			atualizardDivida = Divida{Id: divida.Id}
			o.Read(&atualizardDivida)
		}
	}

	return &atualizardDivida
}

func DeleteDebt(id int) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&Divida{Id: id})
	if err == nil {
		return true
	}

	return false
}

func GetDebt(id int) *Divida {
	o := orm.NewOrm()
	divida := Divida{Id: id}
	o.Read(&divida)
	return &divida
}
