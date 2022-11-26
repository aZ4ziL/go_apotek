package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	DrugID      uint           `json:"drug_id"`
	UserID      uint           `json:"user_id"`
	TotalGoods  uint           `json:"total_goods"`
	TotalPay    uint           `json:"total_pay"`
	TotalRefund uint           `json:"total_refund"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
