package main

import (
	"text/template"

	"github.com/gin-contrib/multitemplate"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFilesFuncs(
		"login",
		template.FuncMap{},
		"views/users/login.tmpl",
	)
	r.AddFromFilesFuncs(
		"register",
		template.FuncMap{},
		"views/users/register.tmpl",
	)

	// Index
	r.AddFromFiles("index", "views/drug/base.tmpl", "views/drug/index.tmpl")

	// Drug
	r.AddFromFilesFuncs(
		"get_all_drugs",
		template.FuncMap{},
		"views/drug/base.tmpl",
		"views/drug/get_all_drugs.tmpl",
	)

	// Transactions
	r.AddFromFilesFuncs(
		"get_all_transactions",
		template.FuncMap{
			"getNameDrugByCode": getNameDrugByCode,
		},
		"views/drug/base.tmpl",
		"views/drug/get_all_transactions.tmpl",
	)
	r.AddFromFilesFuncs(
		"create_transactions",
		template.FuncMap{},
		"views/drug/base.tmpl",
		"views/drug/create_transactions.tmpl",
	)

	return r
}
