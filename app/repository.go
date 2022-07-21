package app

import "core-banking/models"

// Repository ...
type Repository interface {
	GetTransaction(userID uint) (models.UserTransaction, error)
	CreateTransaction(transaction *models.UserTransaction) (*models.UserTransaction, error)
	CreateTransactionHistory(transactionHistory *models.TransactionHistory) (*models.TransactionHistory, error)
	GetTransactionHistory(transHisSearchFilter models.TransHistorySearchFilter) (*[]models.TransactionHistory, error)
}
