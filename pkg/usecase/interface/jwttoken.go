package interfaceUseCase

type IJwtTokenUseCase interface {
	GetDataForCreteAccessToken(string) (string, error)
	GetStatusOfUser(string) (string, error)
}
