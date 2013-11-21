package controllers

import (
	"encoding/json"
	"github.com/robfig/revel"
	"github.com/wlsailor/api-simulator/app/models"
)

type ApiSimulator struct {
	App
}

func (as ApiSimulator) Simulate(path string, productId string) revel.Result {
	as.Validation.Required(path)
	as.Validation.Required(productId)
	if as.Validation.HasErrors() {
		return as.RenderJson(as.Validation.Errors)
	}
	var api *models.ProductApi = &models.ProductApi{}
	err := as.Txn.SelectOne(api, "select * from product_api where product_id = ? and path = ?", productId, "/"+path+"/")
	if err != nil {
		//panic(err)
		return as.RenderJson(err)
	}
	var output interface{}
	data := [] byte(api.Output) 
	err = json.Unmarshal(data, &output)
	if err != nil {
		return as.RenderJson(err)
	}
	return as.RenderJson(output)
}
