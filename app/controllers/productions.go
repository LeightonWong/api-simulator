package	controllers
import	(
	"github.com/robfig/revel"
	"github.com/wlsailor/api-simulator/app/models"
	"string"
	"fmt"
)

type Productions struct {
	App
}

func (p Productions) Index() revel.Result{
	results, err := p.Txn.Select(models.Production{}, `select * from production where ProductionId = ?`, p.connected().ProductionId)
}
