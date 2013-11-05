package	controllers
import	(
	"github.com/robfig/revel"
	"github.com/wlsailor/api-simulator/app/models"
)

type Products struct {
	App
}

func (p Products) Index() revel.Result{
	results, err := p.Txn.Select(models.Product{}, `select * from product`)
	if(err != nil){
		panic(err)
	}
	var products []*models.Product
	for _, r := range results {
		p := r.(*models.Product)
		products = append(products, p)
	}
	return p.Render(products)
}
