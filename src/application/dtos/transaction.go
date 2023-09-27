package dtos

import (
	"time"

	pb "github.com/SardorSaidov8084/restaurant-payment-service/src/application/protos/restaurant_payment"
	"github.com/SardorSaidov8084/restaurant-payment-service/src/domain/payment/models"
)

func ToTransaction(tx *pb.Transaction) *models.Transaction {

	return &models.Transaction{
		ID: tx.GetId(),
		Currency: tx.GetCurrency(),
		Status: tx.GetStatus(),
		CreateTime: int(tx.GetCreateTime()),
		PayTime: int(tx.GetPayTime()),
		CancelTime: int(tx.GetCancelTime()),
		Amount: int(tx.GetAmount()),
		Account: tx.GetAccount(),
		CreatedAt: toTime(tx.GetCreatedAt()),
		UpdatedAt: toTime(tx.GetUpdatedAt()),
	}
}
func ToTransactionPB(tx *models.Transaction) *pb.Transaction {

	return &pb.Transaction{
		Id: tx.ID,
		Currency: tx.Currency,
		Status: tx.Status,
		CreateTime: int32(tx.CreateTime),
		PayTime: int32(tx.PayTime),
		CancelTime: int32(tx.CancelTime),
		Amount: int64(tx.Amount),
		Account: tx.Account,
		CreatedAt: tx.CreatedAt.Format(time.RFC3339),
		UpdatedAt: tx.UpdatedAt.Format(time.RFC3339),
	}
}