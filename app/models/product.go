package	models

import	(
	"github.com/robfig/revel"
	"time"
)

type Product struct {
	ProductId	int	`db:"id"`
	Name	string	`db:"name"`
	Description	string	`db:"description"`
	CreateAt	time.Time	`db:"create_at"`
}

func (p Product) Validate (v * revel.Validation){
	v.Required(p.Name)
	v.Required(p.CreateAt)
}

