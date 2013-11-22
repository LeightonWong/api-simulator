package controllers

import (
	"encoding/json"
	"github.com/robfig/revel"
	"github.com/wlsailor/api-simulator/app/models"
	"log"
	"strings"
)

type ApiSimulator struct {
	App
}

func (as ApiSimulator) Simulate(productId string) revel.Result {
	//as.Validation.Required(path)
	as.Validation.Required(productId)
	if as.Validation.HasErrors() {
		return as.RenderJson(as.Validation.Errors)
	}
	uri := as.Request.RequestURI
	parts := strings.SplitAfterN(uri, "/path", 2)
	if len(parts) != 2 {
		return as.RenderJson("get simulate path error")
	}
	log.Println("simulate path:", parts[1])
	var api *models.ProductApi = &models.ProductApi{}
	err := as.Txn.SelectOne(api, "select * from product_api where product_id = ? and path = ?", productId, parts[1])
	if err != nil {
		//panic(err)
		return as.RenderJson(err)
	}
	var output interface{}
	data := []byte(api.Output)
	err = json.Unmarshal(data, &output)
	if err != nil {
		return as.RenderJson(err)
	}
	return as.RenderJson(output)
}
