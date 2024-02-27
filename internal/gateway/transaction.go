package gateway

import "github.com/flaviomdutra/ms-wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
