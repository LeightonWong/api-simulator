package controllers

import (
	"github.com/robfig/revel"
	"github.com/wlsailor/api-simulator/app/models"
	"time"
)

type ProductApis struct {
	App
}

func (pa ProductApis) List(productId, page, size int) revel.Result {
	pa.Validation.Required(productId)
	if page <= 0 {
		page = 1
	}
	if size <= 0 || size >= 50 {
		size = 20
	}

	results, err := pa.Txn.Select(models.ProductApi{}, "Select * from product_api where product_id = ? limit ?, ?", productId, (page-1)*size, size)
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
	previousPage := page - 1
	nextPage := page + 1
	if previousPage <= 0 {
		previousPage = 1
	}

	if nextPage*size >= int(total) {
		nextPage = page
	}
	return pa.Render(list, productId, total, previousPage, nextPage)
}

func (pa ProductApis) New(productId int) revel.Result {
	return pa.Render(productId)
}

func (pa ProductApis) Add(productId, style int, path, description, input, output string) revel.Result {
	pa.Validation.Required(productId)
	pa.Validation.Required(path)
	api := &models.ProductApi{0, productId, path, description, input, output, style, time.Now()}
	err := pa.Txn.Insert(api)
	if err != nil {
		panic(err)
	}
	pa.Flash.Success("Save new api succeed!")
	return pa.Redirect("/products/%d/apis/list?page=%d&size=%d", productId, 1, 10)
}

func (pa ProductApis) Edit(productId, apiId, style int, path, description, input, output string) revel.Result {
	pa.Validation.Required(productId)
	pa.Validation.Required(apiId)
	pa.Validation.Required(path)
	pa.Validation.Required(style)
	_, err := pa.Txn.Exec("update product_api set path=?, description=?, type=?, input=?, output=? where id = ?", path, description, style, input, output, apiId)
	if err != nil {
		panic(err)
	}
	pa.Flash.Success("Success update api %d", apiId)
	return pa.Redirect("/products/%d/apis/list?page=%d&size=%d", productId, 1, 10)
}

func (pa ProductApis) Detail(productId, apiId int) revel.Result {
	pa.Validation.Required(productId)
	pa.Validation.Required(apiId)
	var api models.ProductApi
	err := pa.Txn.SelectOne(&api, "select * from product_api where id = ?", apiId)
	if err != nil {
		panic(err)
	}
	return pa.Render(api, productId, apiId)
}

func (pa ProductApis) Del(productId, apiId int) revel.Result {
	pa.Validation.Required(apiId)
	_, err := pa.Txn.Exec("delete from product_api where id = ?", apiId)
	if err != nil {
		panic(err)
	}
	pa.Flash.Success("Delete api %d succeeed", apiId)
	return pa.Redirect("/products/%d/apis/list?page=%d&size=%d", productId, 1, 10)
}
