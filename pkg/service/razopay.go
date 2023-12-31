package service

import (
	"fmt"

	"github.com/razorpay/razorpay-go"
	"github.com/razorpay/razorpay-go/utils"
)

func Razopay(totalPrice uint, razopaykey, razopaysecret string) (string, error) {
	client := razorpay.NewClient(razopaykey, razopaysecret)

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

func VerifyPayment(orderID, paymentID, signature, razopaySecret string) bool {

	params := map[string]interface{}{
		"razorpay_order_id":   orderID,
		"razorpay_payment_id": paymentID,
	}

	result := utils.VerifyPaymentSignature(params, signature, razopaySecret)
	fmt.Println("*****", result)
	return result
}
