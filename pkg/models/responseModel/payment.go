package responsemodel

import "time"

type UserWallet struct {
	UserID   string `json:"userID"`
	WalletID string `json:"walletID"`
	Balance  uint   `json:"currentBalance" gorm:"column:balance"`
}

type WalletTransaction struct {
	TransactionID uint      `json:"transactionID"`
	UserID        string    `json:"userID"`
	Credit        uint      `json:"credit,omitempty"`
	Debit         uint      `json:"debit,omitempty"`
	EventDate     time.Time `json:"eventDate"`
	TotalAmount   uint      `json:"totalAmount"`
}
