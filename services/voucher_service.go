package services

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"trinity_app/models"
	"trinity_app/utils"

	"github.com/google/uuid"
)

func (s *services) GenerationVoucher(ctx context.Context, voucher *models.Voucher) (*models.Voucher, error, int) {

	campaign, err, code := s.CheckEligibilityCampaign(ctx, voucher.CampaignID)
	if err != nil {
		return nil, err, code
	}

	voucher.Code = uuid.New().String()

	voucher.ExpiredAt = time.Now().Add(
		utils.Viper.GetDuration("EXPIRED_DURATION"))

	voucher.DiscountPercentage = campaign.DiscountPercentage
	voucher.IsActive = true

	voucher, err = s.Repo.CreateVoucher(ctx, voucher)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	return voucher, nil, http.StatusCreated
}

func (s *services) CheckValidVoucher(ctx context.Context, voucherCode string) (*models.Voucher, error, int) {
	v, err := s.Repo.GetVoucher(ctx, voucherCode)
	if err != nil {
		return nil, err, http.StatusNotFound
	}

	_, err, code := s.CheckEligibilityCampaign(ctx, v.CampaignID)
	if err != nil {
		return nil, fmt.Errorf("[s]Campaign is not valid, %w", err), code
	}

	if time.Now().After(v.ExpiredAt) || !v.IsActive {
		return nil, fmt.Errorf("[s]voucher is not valid"), http.StatusBadRequest
	}

	if v.UserID != nil {
		return nil, fmt.Errorf("[s]voucher is used"), http.StatusBadRequest
	}

	return v, nil, http.StatusOK
}
