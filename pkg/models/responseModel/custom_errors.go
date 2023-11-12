package responsemodel

import "errors"

var (
	IDParamsEmpty        = "id parameter is empty"
	ConversionOFPageErr  = errors.New("attempt to convert string to int made error, page")
	ConversionOfLimitErr = errors.New("attempt to convert string to int made error, page limit")
	PaginationError      = errors.New("page must start from one")
)
