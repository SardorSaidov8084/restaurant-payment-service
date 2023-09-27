package dtos

import (
	"time"

	pb "github.com/SardorSaidov8084/restaurant-payment-service/src/application/protos/restaurant_payment"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/merchant/models"
)

func ToMerchantPB( merchant *models.MerchantSetting) *pb.MerchantSetting {
	return &pb.MerchantSetting{
		EntityId: merchant.EntityID,
		CashboxId: merchant.CashboxID,
		Enabled: merchant.Enabled,
		CreatedAt: merchant.CreatedAt.Format(time.RFC3339),
		UpdatedAt: merchant.UpdatedAt.Format(time.RFC3339),
	}
}