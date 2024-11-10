package repository

import (
	"context"
	"fmt"
	"trinity_app/models"
)

func (r *repository) GetSupscription(ctx context.Context, supscription uint32) (*models.SubscriptionPlanDetails, error) {
	var data models.SubscriptionPlanDetails
	if err := r.db.Where("id = ?", supscription).Take(&data).Error; err != nil {
		return nil, fmt.Errorf("[db]voucher not found")
	}
	return &data, nil
}
