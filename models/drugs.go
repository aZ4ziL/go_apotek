package models

// This is the configuration for the drug model
type Drug struct {
	Code         string        `gorm:"size:10;primaryKey" json:"code"`
	Name         string        `gorm:"size:30" json:"name"`
	Price        uint          `json:"price"`
	Stock        uint          `json:"stock"`
	Transactions []Transaction `gorm:"foreignKey:DrugID" json:"transactions"`
}

// CreateNewDrug
// create new drug
func CreateNewDrug(drug *Drug) error {
	return db.Create(drug).Error
}

// GetAllGrugs
// return all of data by drug
func GetAllGrugs() []Drug {
	var drugs []Drug
	db.Model(&Drug{}).Preload("Transactions").Find(&drugs)
	return drugs
}

// GetDrugByCode
// return drug by passing `code`
func GetDrugByCode(code string) (Drug, error) {
	var drug Drug
	err := db.Model(&Drug{}).Where("code = ?", code).Preload("Transactions").First(&drug).Error
	return drug, err
}
