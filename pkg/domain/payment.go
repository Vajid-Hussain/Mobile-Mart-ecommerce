package domain

type Wallet struct {
	WalletID string `gorm:"primarykey"`
	UserID   uint
	User     Users `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Balance  uint
}
