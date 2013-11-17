package models

import (
	"github.com/robfig/revel"
	"time"
)

type ProductApi struct {
	Id         int       `db:"id"`
	ProductId  int       `db:"product_id"`
	Path       string    `db:"path"`
	Input      string    `db:"input"`
	Output     string    `db:"output"`
	Style      String    `db:"type"`
	UpdateTime time.Time `db:"update_time"`
}

func (pa ProductApi) Validate(v *revel.Validation) {
	v.Required(pa.ProductId)
	v.Required(pa.Path)
}
