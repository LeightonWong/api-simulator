package controllers

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	r "github.com/robfig/revel"
	"github.com/robfig/revel/modules/db/app"
	"github.com/wlsailor/api-simulator/app/models"
	"time"
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

	t := Dbm.AddTable(models.Production{}).SetKeys(true, "ProductionId")
	//t.ColMap("Password").Transient = true
	setColumnSizes(t, map[string]int{
		"Name":     100,
		"CreateAt":	32, 
	})

	Dbm.TraceOn("[gorp]", r.INFO)
	Dbm.CreateTables()

	productions := []*models.Production{
		&models.Production{0, "kmsocial", time.Now()}, 
	}

	if err := Dbm.Insert(productions); err != nil {
		panic(err)
	}

	for _, production := range productions {
		if err := Dbm.Insert(production); err != nil {
			panic(err)
		}
	}
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
