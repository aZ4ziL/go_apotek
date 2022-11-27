package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	DrugCode    string         `json:"drug_id"`
	UserID      uint           `json:"user_id"`
	TotalGoods  uint           `json:"total_goods"`
	TotalPay    uint           `json:"total_pay"`
	TotalPrice  uint           `json:"total_price"`
	TotalRefund uint           `json:"total_refund"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// CreateNewTransaction
// create new transaction
func CreateNewTransaction(trans *Transaction) error {
	return db.Create(trans).Error
}

// GetAllTransactionByUserID
// return all transaction by querying the `UserID`
func GetAllTransactionByUserID(userID uint) []Transaction {
	var transactions []Transaction
	db.Model(&Transaction{}).Where("user_id = ?", userID).Find(&transactions)
	return transactions
}
