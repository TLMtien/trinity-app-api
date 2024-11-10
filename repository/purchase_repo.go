package repository

import (
	"context"
	"fmt"
	"trinity_app/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *repository) StorePurchase(ctx context.Context, p *models.Purchase) (*models.Purchase, error) {

	if err := r.db.Transaction(func(tx *gorm.DB) error {

		var voucher models.Voucher
		if err := tx.Model(&models.Voucher{}).Where("code = ?", p.VoucherCode).
			Update("user_id", p.UserID). // Correct Update syntax
			Take(&voucher).
			Error; err != nil {

			return fmt.Errorf("[db]error update voucher: %w", err)
		}

		if err := tx.Model(&models.Campaign{}).
			Where("id = ?", voucher.CampaignID).
			Update("hold_vouchers", gorm.Expr("hold_vouchers - ?", 1)).Error; err != nil {

			return fmt.Errorf("[db] cannot update campaign: %w", err)
		}

		if err := tx.Create(&p).Clauses(clause.Returning{}).Error; err != nil {
			return fmt.Errorf("[db]cannot store purchase: %w", err)
		}

		if err := tx.Preload("Voucher").
			Preload("SubscriptionPlanPriceDetails").
			Preload("SubscriptionPlanPriceDetails.SubscriptionPlan").
			First(&p, p.ID).Error; err != nil {
			return fmt.Errorf("[db] cannot preload related objects: %w", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return p, nil
}
