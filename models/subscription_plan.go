package models

import (
	"time"
)

type SubscriptionPlan struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	Features  string    `json:"features"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SubscriptionPlanDetails struct {
	ID                 int64            `json:"id"`
	SubscriptionPlanID int64            `json:"subscription_plan_id"`
	SubscriptionPlan   SubscriptionPlan `json:"subscription_plan" gorm:"foreignKey:SubscriptionPlanID"`
	Currency           string           `json:"currency"`
	Price              float64          `json:"price"`
	Plan               string           `json:"plan"`
	Fee                float64          `json:"fee"`
}

func (SubscriptionPlanDetails) TableName() string {
	return "subscription_plan_details"
}

func (SubscriptionPlan) TableName() string {
	return "subscription_plan"
}
