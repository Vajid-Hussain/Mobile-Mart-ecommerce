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

func Pagination(page string, limit string) (int, int, error) {

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

func Validation(data interface{}) (*[]responsemodel.Errors, error) {
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
					err := fmt.Sprintf("%s should be at least %s", e.Field(), e.Param())
					result = responsemodel.Errors{Err: err}
				case "max":
					err := fmt.Sprintf("%s should be at most %s", e.Field(), e.Param())
					result = responsemodel.Errors{Err: err}
				case "email":
					err := fmt.Sprintf("%s should be email structure %s ", e.Field(), e.Param())
					result = responsemodel.Errors{Err: err}
				case "eqfield":
					err := fmt.Sprintf("%s should be equal with %s ", e.Field(), e.Param())
					result = responsemodel.Errors{Err: err}
				case "len":
					err := fmt.Sprintf("%s should be have  %s ", e.Field(), e.Param())
					result = responsemodel.Errors{Err: err}
				case "alpha":
					err := fmt.Sprintf("%s should be Alphabet ", e.Field())
					result = responsemodel.Errors{Err: err}
				case "number":
					err := fmt.Sprintf("%s should be numeric %s ", e.Field(), e.Param())
					result = responsemodel.Errors{Err: err}
				case "numeric":
					err := fmt.Sprintf("%s should be  numeric %s ", e.Field(), e.Param())
					result = responsemodel.Errors{Err: err}
				case "uppercase":
					err := fmt.Sprintf("%s should be  %s %s ", e.Field(), e.Tag(), e.Param())
					result = responsemodel.Errors{Err: err}
				}

				afterErrorCorection = append(afterErrorCorection, result)
			}
		}
		return &afterErrorCorection, errors.New("doesn't fulfill the requirements")
	}
	return &afterErrorCorection, nil
}

func GenerateUUID() string {
	newUUID := uuid.New()

	uuidString := newUUID.String()
	return uuidString
}

func StringToUintConvertion(id string) (uint, error) {

	result, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("converition lead to error")
	}
	return uint(result), nil
}
