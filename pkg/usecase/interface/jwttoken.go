package interfaceUseCase

type IJwtTokenUseCase interface {
	ValidateJwtToken(string) (string, error)
}
