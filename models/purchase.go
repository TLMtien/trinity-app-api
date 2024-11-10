package models

type Purchase struct {
	ID                             int64                   `json:"id"`
	UserID                         int64                   `json:"user_id"`
	VoucherCode                    string                  `json:"voucher_code"`
	Voucher                        Voucher                 `json:"voucher" gorm:"foreignKey:VoucherCode;references:Code"`
	SubscriptionPlanPriceDetailsID int64                   `json:"subscription_plan_details_id"`
	SubscriptionPlanPriceDetails   SubscriptionPlanDetails `json:"subscription_plan_details" gorm:"foreignKey:SubscriptionPlanPriceDetailsID"`
	TotalPrice                     float64                 `json:"total_price"`
}

func (Purchase) TableName() string {
	return "purchases"
}
