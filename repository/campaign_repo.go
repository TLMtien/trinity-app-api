package repository

import (
	"context"
	"fmt"
	"trinity_app/models"
)

func (r *repository) GetCampaign(ctx context.Context, campaignID uint32) (*models.Campaign, error) {

	var campaign models.Campaign

	if err := r.db.Where("id = ?", campaignID).
		Take(&campaign).Error; err != nil {

		return nil, fmt.Errorf("[db] campaign not found: %w", err)
	}
	return &campaign, nil
}
