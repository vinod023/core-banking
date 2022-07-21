package usecase

import (
	"core-banking/app"
	"core-banking/models"
	"errors"
	"time"
)

// UsecaseCoreBanking ...
type UsecaseCoreBanking struct {
	rCharidyCoreBanking app.Repository
}

// NewUsecase ...
func NewUsecase(rCharidyCoreBanking app.Repository) app.Usecase {

	return &UsecaseCoreBanking{
		rCharidyCoreBanking: rCharidyCoreBanking,
	}

}

// UserTransaction - Which creates user transaction
func (u *UsecaseCoreBanking) UserTransaction(transaction models.Transaction) (*models.TransactionHistory, error) {

	userTransaction, err := validateTransaction(u, transaction)
	if err != nil {
		return nil, err
	}

	userTransactionResp, err := u.rCharidyCoreBanking.CreateTransaction(userTransaction)
	if err != nil {
		return nil, err
	}

	// Transfer amount to the other user
	if transaction.TransactionType == models.TRANSFER {
		userTransactionData, err := u.rCharidyCoreBanking.GetTransaction(transaction.UserID)
		if err != nil && err.Error() == "record not found" {
			return nil, errors.New("ivalid transUserId")
		}
		userTransactionData.Balance = userTransactionData.Balance + transaction.Amount

		_, err = u.rCharidyCoreBanking.CreateTransaction(&userTransactionData)
		if err != nil {
			return nil, err
		}
	}

	// TransactionHistory model
	transactionHistory := &models.TransactionHistory{
		UserID:      userTransactionResp.UserID,
		Type:        transaction.TransactionType,
		Amount:      transaction.Amount,
		Balance:     userTransactionResp.Balance,
		TransUserID: transaction.TransUserID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	transactionHistory, err = u.rCharidyCoreBanking.CreateTransactionHistory(transactionHistory)
	if err != nil {
		return nil, err
	}

	return transactionHistory, nil
}

func validateTransaction(u *UsecaseCoreBanking, transaction models.Transaction) (*models.UserTransaction, error) {

	if transaction.UserID == 0 {
		return nil, errors.New("ivalid userId")
	} else if transaction.TransactionType == "" {
		return nil, errors.New("ivalid transactionType")
	}

	userTransaction, err := u.rCharidyCoreBanking.GetTransaction(transaction.UserID)
	if err != nil && err.Error() != "record not found" {
		return nil, err

	} else if transaction.TransUserID == 0 && transaction.TransactionType == models.TRANSFER {
		return nil, errors.New("ivalid transUserId")

	} else if transaction.TransactionType == models.TRANSFER {

		_, err := u.rCharidyCoreBanking.GetTransaction(transaction.TransUserID)
		if err != nil && err.Error() == "record not found" {
			return nil, errors.New("ivalid transUserId")
		}

		balance := userTransaction.Balance - transaction.Amount
		if balance < 0 {
			return nil, errors.New("user doesn't have sufficent balance")
		}
		userTransaction.Balance = balance

	} else if userTransaction.ID == 0 && (transaction.TransactionType == models.WITHDRAW || transaction.TransactionType == models.TRANSFER) {
		return nil, errors.New("invalid user")

	} else if userTransaction.ID == 0 {

		userTransaction.UserID = transaction.UserID
		userTransaction.Balance = transaction.Amount
		userTransaction.CreatedAt = time.Now()

	} else if transaction.TransactionType == models.DEPOSIT {

		userTransaction.Balance = userTransaction.Balance + transaction.Amount

	} else if transaction.TransactionType == models.WITHDRAW {
		balance := userTransaction.Balance - transaction.Amount
		if balance < 0 {
			return nil, errors.New("user doesn't have sufficent balance")
		}
		userTransaction.Balance = balance

	}

	userTransaction.UpdatedAt = time.Now()
	return &userTransaction, nil
}

// GetTransactionHistory - Which returns transactions list based on search cretiria
func (u *UsecaseCoreBanking) GetTransactionHistory(transHisSearchFilter models.TransHistorySearchFilter) (*[]models.TransactionHistory, error) {

	transactionHistory, err := u.rCharidyCoreBanking.GetTransactionHistory(transHisSearchFilter)
	if err != nil {
		return nil, err
	}

	return transactionHistory, nil
}
