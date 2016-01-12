package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/leightonwong/api-simulator/app/models"
	"strings"
)

type ApiSimulator struct {
	App
}

type ErrorMsg struct {
	Errno int      `json:"errno"`
	Err   string   `json:"err"`
	Rst   []string `json:"rst"`
}

func (as ApiSimulator) Simulate(productId string) revel.Result {
	//as.Validation.Required(path)
	as.Validation.Required(productId)
	if as.Validation.HasErrors() {
		return as.RenderJson(ErrorMsg{-2, "parameter error", []string{"you need to specify product id"}})
	}
	uri := as.Request.RequestURI
	parts := strings.SplitAfterN(uri, "/path/", 2)
	if len(parts) != 2 {
		return as.RenderJson("get simulate path error")
	}
	params := as.Params
	url := strings.Split(parts[1], "?")[0]
	completePath := strings.SplitN(url, "/", 2)
	category := completePath[0]
	path := "/" + completePath[1]
	revel.INFO.Println("simulate path:", path, " category:", category)
	var api *models.ProductApi = &models.ProductApi{}
	err := as.Txn.SelectOne(api, "select * from product_api where product_id = ? and category = ? and path = ?", productId, category, path)
	if err != nil {
		panic(err)
		//return as.RenderJson(err)
	}
	errmsgs := []string{}
	if api.Path == "" {
		return as.RenderJson(&ErrorMsg{-1, "No such api , please check the url.", errmsgs})
	}
	//revel.INFO.Println("params", api.Input)
	apiParams := make([]models.ApiParam, 1)
	err = json.Unmarshal([]byte(api.Input), &apiParams)
	if err != nil {
		panic(err)
	}
	for _, param := range apiParams {
		if param.Required == "true" && params.Get(param.Name) == "" {
			errmsgs = append(errmsgs, "required:"+param.Name+" type:"+param.Type)
		}
	}
	if len(errmsgs) > 0 {
		return as.RenderJson(&ErrorMsg{-2, "parameter error", errmsgs})
	}
	var data models.ApiData
	err = as.Txn.SelectOne(&data, "select * from api_data where api_id = ? and valid = true order by rand() limit 0, 1", api.Id)
	if err != nil {
		panic(err)
	}
	var text string
	if data.Output != "" {
		text = data.Output
	} else {
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
