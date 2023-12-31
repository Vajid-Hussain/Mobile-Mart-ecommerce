package resCustomError

import "errors"

var (
	IDParamsEmpty            = "id parameter is empty"
	BindingConflict          = "don't meat data cryteria"
	NotGetSellerIDinContexr  = "not get seller id "
	NotGetUserIdInContexr    = "not get user id "
	ErrConversionOFPage      = errors.New("attempt to convert string to int made error, page")
	ErrConversionOfLimit     = errors.New("attempt to convert string to int made error, page limit")
	ErrPagination            = errors.New("page must start from one")
	ErrPageLimit             = errors.New("page limit must graterthen one")
	ErrNegativeID            = errors.New("id must be grater than one")
	ErrNoRowAffected         = errors.New("no data matching the specified criteria was found in the database")
	ErrProductOrderCompleted = errors.New("the product order is already completed")
	ErrAdminDashbord         = errors.New("face some issue while admin dashbord ")
)
