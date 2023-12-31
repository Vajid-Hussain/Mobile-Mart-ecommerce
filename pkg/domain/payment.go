package domain

import "time"

type Wallet struct {
	WalletID uint  `gorm:"primarykey"`
	UserID   uint  `gorm:"unique"`
	User     Users `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Balance  uint
}

type WalletTransaction struct {
	TransactionID uint `gorm:"primarykey"`
	UserID        uint
	Credit        uint
	Debit         uint
	EventDate     time.Time
	TotalAmount   uint
}
