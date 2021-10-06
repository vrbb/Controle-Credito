package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Endereco struct {
	Id              int
	Logadouro       string    `orm:"null"`
	Numero          string    `orm:"null"`
	Bairro          string    `orm:"null"`
	Cidade          string    `orm:"null"`
	Estado          string    `orm:"null"`
	Complemento     string    `orm:"null"`
	PontoReferencia string    `orm:"null"`
	DataCriacao     time.Time `orm:"auto_now_add;type(datetime)"`
}

func InserirEndereco(endereco Endereco) *Endereco {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Endereco))

	i, _ := qs.PrepareInsert()

	var e Endereco

	// hash cpf {criptografia}
	// endereco.Cpf, _ = hashPassword(endereco.Cpf)

	// get now datetime
	endereco.DataCriacao = time.Now()

	// Insert
	id, err := i.Insert(&endereco)
	if err == nil {
		// successfully inserted
		e = Endereco{Id: int(id)}
		err := o.Read(&e)
		if err == orm.ErrNoRows {
			return nil
		}
	} else {
		return nil
	}

	return &e
}

func BuscarTodasEnderecos() []*Endereco {
	o := orm.NewOrm()
	var enderecos []*Endereco
	o.QueryTable(new(Endereco)).All(&enderecos)

	return enderecos
}

func AtualizarEndereco(endereco Endereco) *Endereco {
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

func DeletarEndereco(id int) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&Endereco{Id: id})
	if err == nil {
		return true
	}

	return false
}

func BuscarEnderecoPeloId(id int) *Endereco {
	o := orm.NewOrm()
	endereco := Endereco{Id: id}
	o.Read(&endereco)
	return &endereco
}
