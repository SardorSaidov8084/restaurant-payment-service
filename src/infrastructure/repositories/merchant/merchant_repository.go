package repositories

import (
	"context"

	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/merchant/models"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/merchant/repositories"
	"gorm.io/gorm"
)

type merchantRepoImpl struct {
	db *gorm.DB
}

func NewMerchantRepoImpl(db *gorm.DB) repositories.MerchantRepository {
	return &merchantRepoImpl{
		db: db,
	}
}
const tableMerchant = "merchant_settings"

func(r *merchantRepoImpl) SaveMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error {

	result := r.db.WithContext(ctx).Table(tableMerchant).Create(setting)

	return result.Error
}
func(r *merchantRepoImpl) UpdateMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error {

	result := r.db.WithContext(ctx).Table(tableMerchant).Save(setting)

	return result.Error
}
func(r *merchantRepoImpl) GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error) {

	var merchant models.MerchantSetting

	result := r.db.WithContext(ctx).Table(tableMerchant).First(&merchant, "entity_id = ?", entityID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &merchant, nil
}