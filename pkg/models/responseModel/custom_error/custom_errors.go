package resCustomError

import "errors"

var (
	IDParamsEmpty           = "id parameter is empty"
	BindingConflict         = "don't meant cryteria, compair with struct"
	NotGetSellerIDinContexr = "not got seller id "
	ErrConversionOFPage     = errors.New("attempt to convert string to int made error, page")
	ErrConversionOfLimit    = errors.New("attempt to convert string to int made error, page limit")
	ErrPagination           = errors.New("page must start from one")
	ErrPageLimit            = errors.New("page limit must graterthen one")
	ErrNegativeID           = errors.New("id must be grater than one")
	ErrNoRowAffected        = errors.New("there is problem with your data, can't did any db proces ")
)
