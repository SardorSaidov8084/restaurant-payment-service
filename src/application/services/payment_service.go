package services

import (
	"context"
	"errors"
	"log"

	"github.com/SardorSaidov8084/restaurant-payment-service/src/application/dtos"
	pb "github.com/SardorSaidov8084/restaurant-payment-service/src/application/protos/restaurant_payment"
	cardsvc "github.com/SardorSaidov8084/restaurant-payment-service/src/domain/card/services"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/payment/models"
	paymentsvc "github.com/SardorSaidov8084/restaurant-payment-service/src/domain/payment/services"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/infrastructure/rand"
)

type PaymentApplicationService interface {
	CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error)
	PayOrderReceipt(ctx context.Context, orderID string) error
	GetReceipt(ctx context.Context, req *pb.GetReceiptRequest) (*pb.GetReceiptResponse, error)
	SaveTransaction(ctx context.Context, req *pb.SaveTransactionRequest) (*pb.SaveTransactionResponse, error)
	UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error)
	GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error)
}

type paymentAppSvcImpl struct {
	paymentSvc paymentsvc.PaymentService
	cardSvc    cardsvc.PaymentCardService
}

func NewPaymentApplicationService(paymentSvc paymentsvc.PaymentService, cardSvc cardsvc.PaymentCardService) PaymentApplicationService {

	return &paymentAppSvcImpl{
		paymentSvc: paymentSvc,
		cardSvc:    cardSvc,
	}
}
func (s *paymentAppSvcImpl) CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error) {

	receipt, err := s.paymentSvc.CreateReceipt(ctx, orderID, cardID, amount)

	if err != nil {
		return nil, err
	}
	log.Printf("Payme Creating receipt %s\n", receipt.OrderID)

	err = s.paymentSvc.UpdateReceipt(ctx, receipt)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}
func (s *paymentAppSvcImpl) GetReceipt(ctx context.Context, req *pb.GetReceiptRequest) (*pb.GetReceiptResponse, error) {

	var (
		receipt *models.Receipt
		err error
	)

	if req.GetReceiptId() != "" {
		receipt, err = s.paymentSvc.GetReceipt(ctx, req.GetReceiptId())
	}else if req.GetOrderId() == ""{
        receipt, err = s.paymentSvc.GetReceiptByOrder(ctx, req.GetOrderId())
	}
	if err !=nil {
		return nil, err
	}
	return &pb.GetReceiptResponse{
		Receipt: dtos.ToReceiptPB(receipt),
	}, nil
}
func (s *paymentAppSvcImpl) PayOrderReceipt(ctx context.Context, orderID string) error {

	receipt, err := s.paymentSvc.GetReceiptByOrder(ctx, orderID)

	if err != nil {
		return err
	}
	card, err := s.cardSvc.GetCard(ctx, receipt.CardID)

	if err != nil {
		return err
	}
	log.Printf("Paying receipt: orderID %s - cardToken %s\n", receipt.OrderID, card.CardToken)

	return nil
}
func (s *paymentAppSvcImpl) SaveTransaction(ctx context.Context, req *pb.SaveTransactionRequest) (*pb.SaveTransactionResponse, error) {

	if req.GetTransaction() == nil {
		return nil, errors.New("Invalid or empty transaction")
	}

	tx := dtos.ToTransaction(req.GetTransaction())

	tx.ID = rand.UUID()
	err := s.paymentSvc.SaveTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}
	return &pb.SaveTransactionResponse{}, nil
}
func (s *paymentAppSvcImpl) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error) {

	if req.GetTransaction() == nil {
		return nil, errors.New("Invalid or empty transaction")
	}
	if req.GetTransaction().GetId() == "" {
		return nil, errors.New("Invalid or missing id")
	}
	transaction := dtos.ToTransaction(req.GetTransaction())
	err := s.paymentSvc.UpdateTransaction(ctx, transaction)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateTransactionResponse{}, nil
}
func (s *paymentAppSvcImpl) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {

	if req.GetTransactionId() == "" {
		return nil, errors.New("Invalid or missing transaction_id")
	}

	tx, err := s.paymentSvc.GetTransaction(ctx, req.GetTransactionId())

	if err != nil {
		return nil, err
	}

	return &pb.GetTransactionResponse{
		Transaction: dtos.ToTransactionPB(tx),
	}, nil
}
