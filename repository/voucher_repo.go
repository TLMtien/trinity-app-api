package repository

import (
	"context"
	"fmt"
	"trinity_app/models"

	"gorm.io/gorm"
)

// check voucher valid
func (r *repository) GetVoucher(ctx context.Context, voucherCode string) (*models.Voucher, error) {
	var data models.Voucher
	if err := r.db.Where("code = ?", voucherCode).Take(&data).Error; err != nil {
		return nil, fmt.Errorf("[db]voucher not found")
	}
	return &data, nil
}

func (r *repository) CreateVoucher(ctx context.Context, voucher *models.Voucher) (*models.Voucher, error) {

	if err := r.db.Create(voucher).Error; err != nil {
		return nil, fmt.Errorf("[db]cannot store voucher: %w", err)
	}
	// update campain to hold
	if err := r.db.Model(&models.Campaign{}).
		Where("id = ?", voucher.CampaignID).
		Update("hold_vouchers", gorm.Expr("hold_vouchers + ?", 1)).
		Update("available_vouchers", gorm.Expr("available_vouchers - ?", 1)).Error; err != nil {

		return nil, fmt.Errorf("[db] cannot update campaign: %w", err)
	}

	return voucher, nil
}
