package interfaces

type IJwtTokenRepository interface {
	GetSellerStatus(string) (string, error)
	GetUserStatus(string) (string, error)
}
