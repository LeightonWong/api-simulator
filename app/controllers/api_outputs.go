package controllers

import (
	"github.com/robfig/revel"
	"github.com/wlsailor/api-simulator/app/models"
	"time"
)

type ApiOutputs struct {
	App
}

func (ao ApiOutputs) List(apiId, page, size int) revel.Result {
	ao.Validation.Required(apiId)
	if page <= 0 {
		page = 1
	}
	if size <= 0 || size > 50 {
		size = 10
	}
	results, err := ao.Txn.Select(models.ApiData{}, "select * from api_data where api_id = ? limit ?, ?", apiId, (page-1)*size, size)
	if err != nil {
		panic(err)
	}
	var apiOutputs []*models.ApiData
	for _, result := range results {
		output := result.(*models.ApiData)
		apiOutputs = append(apiOutputs, output)
	}
	total, err := ao.Txn.SelectInt("select count(*) from api_data where api_id = ?", apiId)
	if err != nil {
		panic(err)
	}
	previousPage := page - 1
	nextPage := page + 1
	if previousPage <= 0 {
		previousPage = 1
	}
	if page*size >= int(total) {
		nextPage = page
	}

	return ao.Render(apiOutputs, apiId, total, previousPage, nextPage)
}

func (ao ApiOutputs) New(apiId int) revel.Result {
	ao.Validation.Required(apiId)
	var apiData models.ProductApi
	err := ao.Txn.SelectOne(&apiData, "select * from product_api where id = ?", apiId)
	if err != nil {
		panic(err)
	}
	return ao.Render(apiData)
}

func (ao ApiOutputs) Add(apiId int, name, output string) revel.Result {
	ao.Validation.Required(apiId)
	ao.Validation.Required(name)
	apiData := &models.ApiData{0, apiId, name, output, true, time.Now()}
	err := ao.Txn.Insert(apiData)
	if err != nil {
		panic(err)
	}
	ao.Flash.Success("Add a new output for api:%d, name:%s", apiId, name)
	return ao.Redirect("/api/%d/outputs/?page=%d&size=%d", apiId, 1, 10)
}

func (ao ApiOutputs) Edit(apiId, dataId int, name, output string) revel.Result {
	ao.Validation.Required(apiId)
	ao.Validation.Required(name)
	apiData := &models.ApiData{dataId, apiId, name, output, true, time.Now()}
	_, err := ao.Txn.Update(apiData)
	if err != nil {
		panic(err)
	}
	ao.Flash.Success("update dataId:%d success", dataId)
	return ao.Redirect("/api/%d/outputs/?page=%d&size=%d", apiId, 1, 10)
}

func (ao ApiOutputs) Detail(dataId int) revel.Result {
	ao.Validation.Required(dataId)
	var apiData models.ApiData
	err := ao.Txn.SelectOne(&apiData, "select * from api_data where id = ?", dataId)
	if err != nil {
		panic(err)
	}
	return ao.Render(apiData, dataId)
}

func (ao ApiOutputs) Del(apiId, dataId int) revel.Result {
	ao.Validation.Required(dataId)
	_, err := ao.Txn.Exec("delete from api_data where id = ?", dataId)
	if err != nil {
		panic(err)
	}
	return ao.Redirect("/api/%d/outputs/?page=%d&size=%d", apiId, 1, 10)
}

func (ao ApiOutputs) ChangeStatus(apiId, dataId int) revel.Result {
	ao.Validation.Required(apiId)
	ao.Validation.Required(dataId)
	_, err := ao.Txn.Exec("update api_data set valid = not valid where id = ?", dataId)
	if err != nil {
		panic(err)
	}
	return ao.Redirect("/api/%d/outputs/?page=%d&size=%d", apiId, 1, 10)
}
