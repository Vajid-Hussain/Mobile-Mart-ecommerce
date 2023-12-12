package requestmodel

type WalletTransaction struct {
	UserID      string `json:"userID"`
	Credit      uint   `json:"credit"`
	Debit       uint   `json:"debit"`
	EventDate   uint   `json:"eventDate"`
	TotalAmount uint   `json:"totalAmount"`
}
