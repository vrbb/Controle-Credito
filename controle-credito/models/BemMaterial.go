package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type BemMaterial struct {
	Id          int
	Descricao   string    `orm:"null"`
	Valor       float32   `orm:"null"`
	DataCriacao time.Time `orm:"auto_now_add;type(datetime)"`
}

func InserirBemMaterial(bemMateiral BemMaterial) *BemMaterial {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BemMaterial))

	i, _ := qs.PrepareInsert()

	var bm BemMaterial

	// hash cpf {criptografia}
	// bemMateiral.Cpf, _ = hashPassword(bemMateiral.Cpf)

	// get now datetime
	bemMateiral.DataCriacao = time.Now()

	// Insert
	id, err := i.Insert(&bemMateiral)
	if err == nil {
		// successfully inserted
		bm = BemMaterial{Id: int(id)}
		err := o.Read(&bm)
		if err == orm.ErrNoRows {
			return nil
		}
	} else {
		return nil
	}

	return &bm
}

func BuscarTodasBemMaterials() []*BemMaterial {
	o := orm.NewOrm()
	var bemMateirals []*BemMaterial
	o.QueryTable(new(BemMaterial)).All(&bemMateirals)

	return bemMateirals
}

func AtualizarBemMaterial(bemMateiral BemMaterial) *BemMaterial {
	o := orm.NewOrm()
	bm := BemMaterial{Id: bemMateiral.Id}
	var atualizardBemMaterial BemMaterial

	// get existing bemMateiral
	if o.Read(&bm) == nil {

		bemMateiral.DataCriacao = bm.DataCriacao
		bm = bemMateiral
		_, err := o.Update(&bm)

		// read updated bemMateiral
		if err == nil {
			// update successful
			atualizardBemMaterial = BemMaterial{Id: bemMateiral.Id}
			o.Read(&atualizardBemMaterial)
		}
	}

	return &atualizardBemMaterial
}

func DeletarBemMaterial(id int) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&BemMaterial{Id: id})
	if err == nil {
		return true
	}

	return false
}

func BuscarBemMaterialPeloId(id int) *BemMaterial {
	o := orm.NewOrm()
	bemMateiral := BemMaterial{Id: id}
	o.Read(&bemMateiral)
	return &bemMateiral
}
