package utils

type RequestGenerateVoucher struct {
	//ID uint32 `json:"id"`
	// UserID     *uint32 `json:"user_id" example:"1"`
	CampaignID uint32 `json:"campaign_id" example:"1"`
}

type RequestPurchase struct {
	UserID                         int64  `json:"user_id"  example:"1"`
	VoucherCode                    string `json:"voucher_code"  example:"0890"`
	SubscriptionPlanPriceDetailsID int64  `json:"subscription_plan_details_id"  example:"1"`
}
