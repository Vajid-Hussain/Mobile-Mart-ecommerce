package responsemodel

type UserWallet struct {
	UserID   string `json:"userID"`
	WalletID string `json:"walletID"`
	Balance  uint   `json:"currentBalance" gorm:"column:balance"`
}
