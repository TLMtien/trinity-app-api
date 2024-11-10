package services

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"trinity_app/models"
)

func (s *services) CheckEligibilityCampaign(ctx context.Context, campaignID uint32) (*models.Campaign, error, int) {
	campaign, err := s.Repo.GetCampaign(ctx, campaignID)
	if err != nil {
		return nil, err, http.StatusNotFound
	}
	timeNow := time.Now()
	if !campaign.IsActive ||
		campaign.StartDate.After(timeNow) ||
		(campaign.EndDate != nil && campaign.EndDate.Before(timeNow)) {
		return nil, fmt.Errorf("[s]campaign is not valid"), http.StatusForbidden
	}

	if campaign.AvailableVouchers <= 0 {
		if campaign.HoldVouchers > 0 {
			return nil, fmt.Errorf("[s]campaign limit reached (holding = %d)", campaign.HoldVouchers), http.StatusForbidden
		}
		return nil, fmt.Errorf("[s]campaign limit reached"), http.StatusForbidden
	}

	return campaign, nil, http.StatusOK
}
