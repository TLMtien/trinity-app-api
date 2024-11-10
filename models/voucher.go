package models

import (
	"time"
)

type Voucher struct {
	ID                 uint32    `json:"id"`
	Code               string    `json:"code"`
	UserID             *uint32   `json:"user_id"`
	CampaignID         uint32    `json:"campaign_id"`
	DiscountPercentage int       `json:"discount_percentage"`
	ExpiredAt          time.Time `json:"expired_at"`
	IsActive           bool      `json:"is_active"`
}

func (Voucher) TableName() string {
	return "vouchers"
}
