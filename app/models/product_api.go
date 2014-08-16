package models

import (
	"github.com/revel/revel"
	"time"
)

type ProductApi struct {
	Id          int       `db:"id"`
	ProductId   int       `db:"product_id"`
	Path        string    `db:"path"`
	Category    string    `db:"category"`
	Description string    `db:"description"`
	Input       string    `db:"input"`
	Output      string    `db:"output"`
	Style       int       `db:"type"`
	UpdateTime  time.Time `db:"update_time"`
}

func (pa ProductApi) Validate(v *revel.Validation) {
	v.Required(pa.ProductId)
	v.Required(pa.Path)
}
