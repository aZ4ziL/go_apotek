package main

import "github.com/aZ4zil/go_apotek/models"

func getNameDrugByCode(code string) string {
	drug, err := models.GetDrugByCode(code)
	if err != nil {
		return ""
	}

	return drug.Name
}
