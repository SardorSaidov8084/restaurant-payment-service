package services

import (
	"context"
	"time"

	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/merchant/models"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/merchant/repositories"
)

type MerchantSettingService interface {
	CreateMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
	UpdateMerchantSetting(ctx context.Context, entityID, cashboxID string, enabled bool) error
	GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
}

type merchantSvcImpl struct {
	merchantRepo repositories.MerchantRepository
}

func NewMerchantService(merchantRepo repositories.MerchantRepository) MerchantSettingService {
	return &merchantSvcImpl{
		merchantRepo: merchantRepo,
	}
}

func (s *merchantSvcImpl) CreateMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error) {

	setting := &models.MerchantSetting{
		EntityID:  entityID,
		CashboxID: "",
		Enabled: false,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := s.merchantRepo.SaveMerchantSetting(ctx, setting)

	if err != nil {
		return nil, err
	}

	return setting, nil
}
func (s *merchantSvcImpl) UpdateMerchantSetting(ctx context.Context, entityID, cashboxID string, enabled bool) error {

	setting, err := s.merchantRepo.GetMerchantSetting(ctx, entityID)

	if err != nil {
		return err
	}
	setting.CashboxID = cashboxID
	setting.Enabled = enabled
	setting.UpdatedAt = time.Now().UTC()
	return s.merchantRepo.UpdateMerchantSetting(ctx, setting)
}
func (s *merchantSvcImpl) GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error) {
	
	return s.merchantRepo.GetMerchantSetting(ctx, entityID)
}
