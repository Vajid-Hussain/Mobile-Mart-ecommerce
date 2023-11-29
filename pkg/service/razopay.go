package service

import (
	"github.com/razorpay/razorpay-go"
)

func Razopay(totalPrice uint) (string, error) {
	client := razorpay.NewClient("rzp_test_TvFtCr7NADxnEC", "qvxbhiwTJZLHHE3tNQQv8Mty")

	data := map[string]interface{}{
		"amount":   totalPrice * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		return "", err
	}
	idFromResponse, _ := body["id"].(string)
	return idFromResponse, nil
}
