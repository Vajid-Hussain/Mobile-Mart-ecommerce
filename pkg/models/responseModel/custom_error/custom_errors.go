package resCustomError

import "errors"

var (
	IDParamsEmpty        = "id parameter is empty"
	BindingConflict      = "can't bind json with struct"
	ErrConversionOFPage  = errors.New("attempt to convert string to int made error, page")
	ErrConversionOfLimit = errors.New("attempt to convert string to int made error, page limit")
	ErrPagination        = errors.New("page must start from one")
	ErrPageLimit         = errors.New("page limit must graterthen one")
)
