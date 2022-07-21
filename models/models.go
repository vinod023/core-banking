package models

import "time"

type (
	User struct {
		UserID    uint      ` gorm:"column:id" json:"userId"`
		FirstName string    `gorm:"column:fname"  json:"firstName"`
		LastName  string    `gorm:"column:lname"  json:"lastName"`
		Email     string    `gorm:"column:email" json:"email"`
		Mobile    string    `gorm:"mobile" json:"mobile"`
		CreatedAt time.Time `gorm:"column:created_at"  json:"createdAt"`
		UpdatedAt time.Time `gorm:"column:updated_at"  json:"updatedAt"`
	}

	UserTransaction struct {
		ID        uint      ` gorm:"column:id" json:"id"`
		UserID    uint      ` gorm:"column:user_id" json:"userId"`
		Balance   float64   ` gorm:"column:balance" json:"balance"`
		CreatedAt time.Time `gorm:"column:created_at"  json:"createdAt"`
		UpdatedAt time.Time `gorm:"column:updated_at"  json:"updatedAt"`
	}

	TransactionHistory struct {
		ID          uint      ` gorm:"column:id" json:"id"`
		UserID      uint      ` gorm:"column:userid" json:"userId"`
		Type        string    ` gorm:"column:type" json:"type"`
		Amount      float64   ` gorm:"column:amount" json:"amount"`
		Balance     float64   ` gorm:"column:balance" json:"balance"`
		TransUserID uint      `gorm:"column:trans_userid" json:"transUserId"`
		CreatedAt   time.Time `gorm:"column:created_at"  json:"createdAt"`
		UpdatedAt   time.Time `gorm:"column:updated_at"  json:"updatedAt"`
	}

	Transaction struct {
		UserID          uint    `json:"userId"`
		TransactionType string  `json:"transactionType"`
		Amount          float64 `json:"amount"`
		TransUserID     uint    `json:"transUserId"`
	}

	TransHistorySearchFilter struct {
		UserID          uint       `json:"userId"`
		Amount          float64    `json:"amount"`
		Balance         float64    `json:"balance"`
		TransactionType string     `json:"transactionType"`
		StartDate       *time.Time `json:"startDate"`
		EndDate         *time.Time `json:"endDate"`
	}

	Response struct {
		Status        string  `json:"status"`
		TransactionID uint    `json:"transactionId"`
		Balance       float64 `json:"balance"`
	}
)

func (*User) TableName() string {
	return "users"
}

func (*UserTransaction) TableName() string {
	return "user_transaction"
}

func (*TransactionHistory) TableName() string {
	return "transaction_history"
}
