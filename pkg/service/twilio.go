package service

import (
	"errors"
	"fmt"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	serviceInterface "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service/interface"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

type twilioOtp struct {
	requirements config.OTP
}

func NewOtpService(details config.OTP) serviceInterface.Ijwt {
	return &twilioOtp{requirements: details}
}

var tw *twilio.RestClient

func (o *twilioOtp) TwilioSetup() {
	tw = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: o.requirements.AccountSid,
		Password: o.requirements.AuthToken,
	})
}

func (o *twilioOtp) SendOtp(phone string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo("+91" + phone)
	params.SetChannel("sms")
	res, err := tw.VerifyV2.CreateVerification(o.requirements.ServiceSid, params)
	if err != nil {
		return "", err
	}
	return *res.Sid, nil
}

func (o *twilioOtp) VerifyOtp(phone string, otp string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(otp)
	res, err := tw.VerifyV2.CreateVerificationCheck(o.requirements.ServiceSid, params)
	fmt.Println("res.status on otp verification", *res.Status)
	if err != nil {
		return errors.New("eroor")
	}

	if *res.Status == "approved" {
		return nil
	}
	return errors.New("failed to verify otp")
}
