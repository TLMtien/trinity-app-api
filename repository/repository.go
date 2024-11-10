package repository

import (
	"context"
	"trinity_app/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetVoucher(ctx context.Context, voucherCode string) (*models.Voucher, error)
	CreateVoucher(ctx context.Context, voucher *models.Voucher) (*models.Voucher, error)

	GetSupscription(ctx context.Context, supscription uint32) (*models.SubscriptionPlanDetails, error)

	StorePurchase(ctx context.Context, p *models.Purchase) (*models.Purchase, error)

	GetCampaign(ctx context.Context, campaignID uint32) (*models.Campaign, error)
}

func NewRepo(gormDB *gorm.DB) Repository {
	return &repository{gormDB}
}
