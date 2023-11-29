package service

import (
	"fmt"

	"github.com/razorpay/razorpay-go"
	"github.com/razorpay/razorpay-go/utils"
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

func VerifyPayment(orderID, paymentID, signature string) bool {

	params := map[string]interface{}{
		"razorpay_order_id":   orderID,
		"razorpay_payment_id": paymentID,
	}

	// signature := "0d4e745a1838664ad6c9c9902212a32d627d68e917290b0ad5f08ff4561bc50f"
	secret := "qvxbhiwTJZLHHE3tNQQv8Mty"
	result := utils.VerifyPaymentSignature(params, signature, secret)
	fmt.Println("*****", result)
	return result
}
