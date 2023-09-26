package services

import (
	"context"
	"time"

	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/merchant/models"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/merchant/repositories"
)

type MerchantSetting interface {
	CreateMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
	UpdateMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error
	GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
}

type merchantSvcImpl struct {
	merchantRepo repositories.MerchantRepository
}
func NewMerchantService( merchantRepo repositories.MerchantRepository) MerchantSetting {
	return &merchantSvcImpl{
		merchantRepo: merchantRepo,
	}
}

func (s *merchantSvcImpl) CreateMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error){

	setting := &models.MerchantSetting{
		EntityID: entityID,
		CashboxID: "",
		Enabled: true,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := s.merchantRepo.SaveMerchantSetting(ctx, setting)

	if err != nil {
		return nil, err
	}

	return setting, nil
}
func (s *merchantSvcImpl) UpdateMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error
func (s *merchantSvcImpl) GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
