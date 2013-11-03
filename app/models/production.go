package	models

import	(
	"fmt"
	"github.com/robfig/revel"
	"regexp"
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
	v.Match(p.Name, regexp.MustCompile('\s')
}

