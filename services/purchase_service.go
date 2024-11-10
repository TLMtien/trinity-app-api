package services

import (
	"context"
	"log"
	"net/http"
	"trinity_app/models"
)

func (s *services) Purchase(ctx context.Context, p *models.Purchase) (*models.Purchase, error, int) {
	log.Print(p)
	sup, err := s.Repo.GetSupscription(ctx, uint32(p.SubscriptionPlanPriceDetailsID))
	if err != nil {
		return nil, err, http.StatusNotFound
	}

	v, err, code := s.CheckValidVoucher(ctx, p.VoucherCode)
	if err != nil {
		return nil, err, code
	}

	p.TotalPrice = sup.Price * (1 - float64(v.DiscountPercentage)/100)

	p, err = s.Repo.StorePurchase(ctx, p)

	if err != nil {
		return nil, err, http.StatusInternalServerError
	}
	return p, nil, http.StatusCreated

}
