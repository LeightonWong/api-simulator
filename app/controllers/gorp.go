package controllers

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	r "github.com/revel/revel"
	"github.com/revel/revel/modules/db/app"
	"github.com/wlsailor/api-simulator/app/models"
	//"time"
)

var (
	Dbm *gorp.DbMap
)

func Init() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.MySQLDialect{}}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTableWithName(models.Product{}, "product").SetKeys(true, "ProductId")
	//t.ColMap("Password").Transient = true
	setColumnSizes(t, map[string]int{
		"Name":        100,
		"Description": 512,
		"CreateAt":    32,
	})

	Dbm.AddTableWithName(models.ProductApi{}, "product_api").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.ApiData{}, "api_data").SetKeys(true, "Id")

	Dbm.TraceOn("[gorp]", r.INFO)
	//Dbm.CreateTablesIfNotExists()

	/*
		products := []*models.Product{
			&models.Product{0, "kmsocial", time.Now()},
			&models.Product{0, "dec", time.Now()},
		}

		for _, product := range products {
			if err := Dbm.Insert(product); err != nil {
				panic(err)
			}
		}
	*/
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
