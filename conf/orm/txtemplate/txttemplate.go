package txtemplate

import (
	"github.com/beego/beego/v2/client/orm"
)

type TxTemplate struct {
	ormer orm.Ormer
}

func NewTxTemplate() *TxTemplate {
	return &TxTemplate{
		ormer: orm.NewOrm(),
	}
}

func (t *TxTemplate) Tx(fn func(o orm.Ormer) (interface{}, error)) (interface{}, error) {
	tx, err := t.ormer.Begin()
	if err != nil {
		return nil, err
	}
	result, err := fn(t.ormer)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return result, nil
}
