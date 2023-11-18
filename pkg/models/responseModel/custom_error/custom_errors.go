package resCustomError

import "errors"

var (
	IDParamsEmpty           = "id parameter is empty"
	BindingConflict         = "don't meant data cryteria"
	NotGetSellerIDinContexr = "not got seller id "
	NotGetUserIdInContexr   = "not got user id "
	ErrConversionOFPage     = errors.New("attempt to convert string to int made error, page")
	ErrConversionOfLimit    = errors.New("attempt to convert string to int made error, page limit")
	ErrPagination           = errors.New("page must start from one")
	ErrPageLimit            = errors.New("page limit must graterthen one")
	ErrNegativeID           = errors.New("id must be grater than one")
	ErrNoRowAffected        = errors.New("equalent data is not exist in db")
)
