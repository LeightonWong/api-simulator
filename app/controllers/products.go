package controllers

import (
	"github.com/robfig/revel"
	"github.com/wlsailor/api-simulator/app/models"
	"time"
)

type Products struct {
	App
}

func (p Products) Index() revel.Result {
	results, err := p.Txn.Select(models.Product{}, `select * from product`)
	if err != nil {
		panic(err)
	}
	var products []*models.Product
	for _, r := range results {
		p := r.(*models.Product)
		products = append(products, p)
	}
	return p.Render(products)
}

func (p Products) List(size, page int) revel.Result {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	var products []*models.Product
	results, err := p.Txn.Select(models.Product{}, `select * from product limit ?, ?`, (page-1)*size, size)
	if err != nil {
		panic(err)
	}
	for _, r := range results {
		p := r.(*models.Product)
		products = append(products, p)
	}
	nextPage := page + 1
	total, err := p.Txn.SelectInt(`select count(*) from product`)
	if err != nil {
		panic(err)
	}
	return p.Render(products, nextPage, total)
}

func (p Products) Add(name, description string) revel.Result {
	p.Validation.Required(name)
	p.Validation.Required(description)
	product := models.Product{Name: name, Description: description, CreateAt: time.Now()}
	err := p.Txn.Insert(&product)
	if err != nil {
		panic(err)
	}
	return p.Redirect(Products.List, 10, 1)
}

func (p Products) Edit(id, name, description string) revel.Result {
	p.Validation.Required(name)
	p.Validation.Required(description)
	_, err := p.Txn.Exec(`update product set name = ? , description = ? where id = ?`, name, description, id)
	if err != nil {
		panic(err)
	}
	return p.Redirect(Products.List, 10, 1)
}

func (p Products) Detail(id string) revel.Result {
	var product *models.Product = &models.Product{}
	err := p.Txn.SelectOne(product, "select * from product where id = ?", id)
	if err != nil {
		panic(err)
	}
	//product := result.(*models.Product)
	return p.Render(product)
}

func (p Products) Del(id string) revel.Result {
	_, err := p.Txn.Exec("delete from product where id = ?", id)
	if err != nil {
		panic(err)
	}
	return p.Redirect(Products.List, 10, 1)
}

func (p Products) New() revel.Result {
	return p.Render()
}
