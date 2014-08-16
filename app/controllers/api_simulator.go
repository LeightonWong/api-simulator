package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
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
		panic(err)
		//return as.RenderJson(err)
	}
	var data models.ApiData
	err = as.Txn.SelectOne(&data, "select * from api_data where api_id = ? order by rand() limit 0, 1", api.Id)
	if err != nil {
		panic(err)
	}
	var text string
	if data.Output != "" {
		text = data.Output
	}else{
		text = api.Output
	}
	var output interface{}
	result := []byte(text)
	err = json.Unmarshal(result, &output)
	if err != nil {
		return as.RenderJson(err)
	}
	return as.RenderJson(output)
}
