package interfaces

type IAdminRepository interface {
	GetPassword(string) (string, error)
}
