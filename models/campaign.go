package models

import "time"

type Campaign struct {
	ID                 int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name               string     `gorm:"not null" json:"name"`
	Description        string     `json:"description"`
	IsActive           bool       `json:"is_active"`
	DiscountPercentage int        `json:"discount_percentage"`
	StartDate          time.Time  `json:"start_date"`
	EndDate            *time.Time `json:"end_date"`
	MaxVouchers        int        `json:"max_vouchers"`
	HoldVouchers       int        `json:"hold_vouchers"`
	AvailableVouchers  int        `json:"available_vouchers"`
}

func (Campaign) TableName() string {
	return "campaigns"
}
