package models

import (
	"github.com/robfig/revle"
	"time"
)

type ApiData struct {
	Id         int       `db:"id"`
	ApiId      int       `db:"api_id"`
	DataType   string    `db:"type"`
	Output     string    `db:"output"`
	UpdateTime time.Time `db:"update_time"`
}

func (ad ApiData) Validate(v *revel.Validation) {
	v.Required(ad.ApiId)
	v.Required(ad.DataType)
}
