package app

import "core-banking/models"

// Usecase ...
type Usecase interface {
	UserTransaction(transaction models.Transaction) (*models.TransactionHistory, error)
	GetTransactionHistory(transHisSearchFilter models.TransHistorySearchFilter) (*[]models.TransactionHistory, error)
}
