package repository

import (
	"core-banking/app"
	"core-banking/models"
	"strings"

	"github.com/jinzhu/gorm"
)

// RepositoryCoreBanking ...
type RepositoryCoreBanking struct {
	db *gorm.DB
}

// NewRepository ...
func NewRepository(db *gorm.DB) app.Repository {

	return &RepositoryCoreBanking{
		db: db,
	}

}

// GetTransaction which retrives transaction based on user
func (r *RepositoryCoreBanking) GetTransaction(userID uint) (transaction models.UserTransaction, err error) {

	if err = r.db.Where("user_id = ?", userID).Find(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

// CreateTransaction create the transaction
func (r *RepositoryCoreBanking) CreateTransaction(transaction *models.UserTransaction) (*models.UserTransaction, error) {

	if err := r.db.Save(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

// CreateTransactionHistory create the transaction history
func (r *RepositoryCoreBanking) CreateTransactionHistory(transactionHistory *models.TransactionHistory) (*models.TransactionHistory, error) {

	if err := r.db.Save(&transactionHistory).Error; err != nil {
		return nil, err
	}

	return transactionHistory, nil
}

// GetTransactionHistory Retrives te transactions history based on search creteria
func (r *RepositoryCoreBanking) GetTransactionHistory(transHisSearchFilter models.TransHistorySearchFilter) (*[]models.TransactionHistory, error) {
	transactionHistory := []models.TransactionHistory{}

	db := r.db

	if transHisSearchFilter.UserID > 0 {
		db = db.Where("userid = ?", transHisSearchFilter.UserID)
	}

	if transHisSearchFilter.Amount > 0 {
		db = db.Where("amount >= ?", transHisSearchFilter.Amount)
	}

	if transHisSearchFilter.Balance > 0 {
		db = db.Where("balance >= ?", transHisSearchFilter.Balance)
	}

	if transHisSearchFilter.TransactionType != "" {

		var trainsArr []string
		if strings.Contains(transHisSearchFilter.TransactionType, ",") {
			trainsArr = strings.Split(transHisSearchFilter.TransactionType, ",")
		} else {
			trainsArr = append(trainsArr, transHisSearchFilter.TransactionType)
		}

		db = db.Where("type in (?) ", trainsArr)
	}

	if transHisSearchFilter.StartDate != nil && !transHisSearchFilter.StartDate.IsZero() {
		db = db.Where("created_at >= ?", transHisSearchFilter.StartDate)
	}

	if transHisSearchFilter.EndDate != nil && !transHisSearchFilter.EndDate.IsZero() {
		db = db.Where("created_at <= ?", transHisSearchFilter.EndDate)
	}

	if err := db.Find(&transactionHistory).Error; err != nil {
		return nil, err
	}

	return &transactionHistory, nil
}
