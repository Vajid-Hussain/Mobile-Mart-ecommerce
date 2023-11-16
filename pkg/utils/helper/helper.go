package helper

import (
	"errors"
	"fmt"
	"strconv"

	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func PaginationError(page string, limit string) (int, int, error) {

	pageNO, err := strconv.Atoi(page)
	if err != nil {
		return 0, 0, resCustomError.ErrConversionOFPage
	}

	if pageNO < 1 {
		return 0, 0, resCustomError.ErrPagination
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		return 0, 0, resCustomError.ErrConversionOfLimit
	}

	if limits <= 0 {
		return 0, 0, resCustomError.ErrPageLimit
	}

	offSet := (pageNO * limits) - limits
	limits = pageNO * limits

	return offSet, limits, nil
}

func validation(data interface{}) (*[]responsemodel.Errors, error) {
	var afterErrorCorection []responsemodel.Errors
	var result responsemodel.Errors
	validate := validator.New()

	err := validate.Struct(data)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Tag() {
				case "required":
					err := fmt.Sprintf("%s is required", e.Field())
					result = responsemodel.Errors{Err: err}
				case "min":
					err := fmt.Sprintf("%s should be at least %s characters", e.Field(), e.Param())
					result = responsemodel.Errors{Err: err}
				case "max":
					err := fmt.Sprintf("%s should be at most %s characters", e.Field(), e.Param())
					result = responsemodel.Errors{Err: err}
				}
				afterErrorCorection = append(afterErrorCorection, result)
			}
		}
		return &afterErrorCorection, errors.New("doesn't fulfill the inventory requirements")
	}
	return &afterErrorCorection, nil
}

func GenerateUUID() string {
	newUUID := uuid.New()

	uuidString := newUUID.String()
	return uuidString
}
