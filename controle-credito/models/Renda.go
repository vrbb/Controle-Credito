package models

import (
	"controle-credito/conf"
	"time"

	"github.com/astaxie/beego/orm"
)

type Renda struct {
	Id          int       `orm:"not_null;unique"`
	Descricao   string    `orm:"null"`
	Valor       float32   `orm:"null"`
	DataCriacao time.Time `orm:"auto_now_add;type(datetime)"`
	Cpf         string    `orm:"not_null;unique"`
	conf.Encryptionkey
}

func AddIncome(renda Renda) *Renda {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Renda))

	i, _ := qs.PrepareInsert()

	var r Renda

	// hash cpf {criptografia}
	// renda.Cpf, _ = hashPassword(renda.Cpf)

	// get now datetime
	renda.DataCriacao = time.Now()

	// Insert
	id, err := i.Insert(&renda)
	if err == nil {
		// successfully inserted
		r = Renda{Id: int(id)}
		err := o.Read(&r)
		if err == orm.ErrNoRows {
			return nil
		}
	} else {
		return nil
	}

	return &r
}

func AllIncome() []*Renda {
	o := orm.NewOrm()
	var rendas []*Renda
	o.QueryTable(new(Renda)).All(&rendas)

	return rendas
}

func UpdateIncome(renda Renda) *Renda {
	o := orm.NewOrm()
	r := Renda{Id: renda.Id}
	var atualizardRenda Renda

	// get existing renda
	if o.Read(&r) == nil {

		renda.DataCriacao = r.DataCriacao
		r = renda
		_, err := o.Update(&r)

		// read updated renda
		if err == nil {
			// update successful
			atualizardRenda = Renda{Id: renda.Id}
			o.Read(&atualizardRenda)
		}
	}

	return &atualizardRenda
}

func DeleteIncome(id int) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&Renda{Id: id})
	if err == nil {
		return true
	}

	return false
}

func GetIncome(id int) *Renda {
	o := orm.NewOrm()
	renda := Renda{Id: id}
	o.Read(&renda)
	return &renda
}
