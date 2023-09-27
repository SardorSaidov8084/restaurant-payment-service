package grpc

import (
	"context"

	pb "github.com/SardorSaidov8084/restaurant-payment-service/src/application/protos/restaurant_payment"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/application/services"
)

type Server struct {
	pb.RestaurantPaymentServer
	merchantAppSvc services.MerchantApplicationService
	paymentAppSvc services.PaymentApplicationService
}

func NewServer(merchantAppSvc services.MerchantApplicationService, paymentAppSvc services.PaymentApplicationService) *Server {
	return &Server{
		merchantAppSvc: merchantAppSvc,
		paymentAppSvc: paymentAppSvc,
	}
}

func (s *Server) UpdateMerchantSetting(ctx context.Context, req *pb.UpdateMerchantSettingRequest) (*pb.UpdateMerchantSettingResponse, error) {
	return s.merchantAppSvc.UpdateMerchantSetting(ctx, req)
}

func (s *Server) GetMerchantSetting(ctx  context.Context, req  *pb.GetMerchantSettingRequest) (*pb.GetMerchantSettingResponse, error) {

	return s.merchantAppSvc.GetMerchantSetting(ctx, req)
}
func (s *Server) GetReceipt(ctx context.Context, req *pb.GetReceiptRequest) (*pb.GetReceiptResponse, error) {

	return s.paymentAppSvc.GetReceipt(ctx, req)
}
func (s *Server) SaveTransaction(ctx context.Context, req *pb.SaveTransactionRequest) (*pb.SaveTransactionResponse, error) {
	return s.paymentAppSvc.SaveTransaction(ctx, req)
}
func (s *Server) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error) {
	return s.paymentAppSvc.UpdateTransaction(ctx, req)
}
func (s *Server) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	return s.paymentAppSvc.GetTransaction(ctx, req)
}
