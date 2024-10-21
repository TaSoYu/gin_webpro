package models

import "time"

// 汇率结构体
type ExchangeRate struct {
	ID           uint      `gorm:"primary_key;auto_increment" json:"_id"`
	FromCurrency string    `json:"fromCurrency" binding:"required"`
	ToCurrency   string    `json:"toCurrency" binding:"required"`
	Rate         float64   `json:"rate" binding:"required"`
	Date         time.Time `json:"date"`
}
