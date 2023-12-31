package interfaces

type IAdminRepository interface {
	GetPassword(string) (string, error)

	GetSellerDetailsForDashBord(string) (uint, error)
	TotalRevenue() (uint, uint, error)
	GetNetCredit() (uint, error)
}
