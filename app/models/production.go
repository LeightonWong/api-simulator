package	models

import	(
	"github.com/robfig/revel"
	"time"
)

type Production struct {
	ProductionId	int
	Name	string
	CreateAt	time.Time
}

func (p Production) Validate (v * revel.Validation){
	v.Required(p.Name)
	v.Required(p.CreateAt)
}

