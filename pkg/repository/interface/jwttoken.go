package interfaces

type IJwtTokenRepository interface {
	GetUserStatus(string) (string, error)
}
