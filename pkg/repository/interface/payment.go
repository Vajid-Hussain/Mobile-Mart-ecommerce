package interfaces

type IPaymentRepository interface {
	CreateOrUpdateWallet(string, uint) (*uint, error)
}
