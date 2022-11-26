package models

// This is the configuration for the drug model
type Drug struct {
	Code         string        `gorm:"size:10;primaryKey" json:"code"`
	Name         string        `gorm:"size:30" json:"name"`
	Price        uint          `json:"price"`
	Stock        uint          `json:"stock"`
	Transactions []Transaction `gorm:"foreignKey:DrugID" json:"transactions"`
}
