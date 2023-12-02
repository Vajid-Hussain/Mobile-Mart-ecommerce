package domain

type Wallet struct {
	WalletID uint  `gorm:"primarykey"`
	UserID   uint  `gorm:"unique"`
	User     Users `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	Balance  uint
}
