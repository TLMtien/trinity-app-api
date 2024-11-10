package services

import (
	"context"
	"trinity_app/models"
	"trinity_app/repository"
)

type services struct {
	Repo repository.Repository
}

type Services interface {
	GenerationVoucher(ctx context.Context, voucher *models.Voucher) (*models.Voucher, error, int)
	CheckValidVoucher(ctx context.Context, voucherCode string) (*models.Voucher, error, int)
	CheckEligibilityCampaign(ctx context.Context, campaignID uint32) (*models.Campaign, error, int)
	Purchase(ctx context.Context, data *models.Purchase) (*models.Purchase, error, int)
}

func NewServices(repo repository.Repository) Services {
	return &services{repo}
}
