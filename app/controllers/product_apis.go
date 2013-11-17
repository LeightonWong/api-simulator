package controllers

import (
	"github.com/robfig/revel"
	"github.com/wlsailor/api_simulator/app/models"
	"time"
)

type ProductApis struct {
	App
}

func (pa ProductApis) List(productId, page, size int) revel.Result {
	revel.Validation.Required(productId)
	if page <= 0 {
		page = 1
	}
	if size <= 0 || size >= 50 {
		size = 20
	}

	results, err := pa.Txn.Select(models.ProductApi{}, "Select * from ProductApi where ProductId = ? limit ?, ?", productId, (page-1)*size, size)
	if err != nil {
		panic(err)
	}
	var list []*models.ProductApi
	for _, r := range results {
		api := r.(*models.ProductApi)
		list = append(list, api)
	}
	total, err := pa.Txn.SelectInt("select count(*) from product_api where product_id = ?", productId)
	if err != nil {
		panic(err)
	}
	nextPage := page + 1
	return p.Render(list, total, nextPage)
}

func (pa ProductApis) New() revel.Result {
	return pa.Render()
}

func (pa ProductApis) Add(productId, style int, path, input, output string) revel.Result {
	revel.Validation.Required(productId)
	revel.Validation.Required(path)
	api := &models.ProductApi{0, productId, path, input, output, style, time.Now()}
	err := pa.Txn.Insert(api)
	if err != nil {
		panic(err)
	}
	return pa.Render(ProductApis.List, 1, 20)
}
