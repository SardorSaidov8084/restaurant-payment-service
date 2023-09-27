package services

import (
	"context"
	"errors"

	"github.com/SardorSaidov8084/restaurant-payment-service/src/application/dtos"
	pb "github.com/SardorSaidov8084/restaurant-payment-service/src/application/protos/restaurant_payment"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/merchant/services"
)
type MerchantApplicationService interface {
	UpdateMerchantSetting(ctx context.Context, req *pb.UpdateMerchantSettingRequest) (*pb.UpdateMerchantSettingResponse, error)
	GetMerchantSetting(ctx  context.Context, req  *pb.GetMerchantSettingRequest) (*pb.GetMerchantSettingResponse, error)
}

type merchantAppSvcImpl struct {
	merchantSvc services.MerchantSettingService
}

func NewMerchantApplicationService(merchantSvc services.MerchantSettingService) MerchantApplicationService{
	return &merchantAppSvcImpl {
		merchantSvc: merchantSvc,
	}
}

func (s *merchantAppSvcImpl) UpdateMerchantSetting(ctx context.Context, req *pb.UpdateMerchantSettingRequest) (*pb.UpdateMerchantSettingResponse, error) {

	if req.GetEntityId() == "" {
		return nil, errors.New("Invalid or missing entity_id")
	}

	err := s.merchantSvc.UpdateMerchantSetting(ctx, req.GetEntityId(), req.GetCashboxId(), req.GetEnabled())

	if err != nil {
		return nil, err
	}

	return &pb.UpdateMerchantSettingResponse{}, nil
}
func (s *merchantAppSvcImpl) GetMerchantSetting(ctx  context.Context, req  *pb.GetMerchantSettingRequest) (*pb.GetMerchantSettingResponse, error) {

	if req.GetEntityId() == "" {
		return nil, errors.New("Invalid or missing entity_id")
	}

	merchant, err := s.merchantSvc.GetMerchantSetting(ctx, req.GetEntityId())

	if err != nil {
		return nil, err
	}

	return &pb.GetMerchantSettingResponse{
		Setting: dtos.ToMerchantPB(merchant),
	}, nil
}