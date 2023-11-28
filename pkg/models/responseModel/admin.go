package responsemodel

type AdminLoginRes struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Result   string `json:"result,omitempty"`
	Token    string `json:"token,omitempty"`
	// RefreshToken string `json:"refreshtoken,omitempty"`
}

type AdminDashBord struct {
	TotalSellers   uint `json:"totalSellers"`
	BlockedSellers uint `json:"blockedSellers"`
	ActiveSellers  uint `json:"activeSellers"`
	TotalRevenue   uint `json:"totalRevenue"`
	TotalOrders    uint `json:"totalOrders"`
	TotalCredit    uint `json:"totalCredit"`
}
