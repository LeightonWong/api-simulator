package models

import (
	"github.com/revel/revel"
	"time"
)

type ApiData struct {
	Id         int       `db:"id"`
	ApiId      int       `db:"api_id"`
	Name       string    `db:"name"`
	Output     string    `db:"output"`
	Valid      bool      `db:valid`
	UpdateTime time.Time `db:"update_time"`
}

func (ad ApiData) Validate(v *revel.Validation) {
	v.Required(ad.ApiId)
	v.Required(ad.Name)
}
