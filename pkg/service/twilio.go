package service

import (
	"errors"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

type twilioOtp struct {
	requirements config.OTP
}

var twilioOTP twilioOtp

func OtpService(details config.OTP) {
	twilioOTP.requirements = details
}

var tw *twilio.RestClient

func TwilioSetup() {
	tw = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilioOTP.requirements.AccountSid,
		Password: twilioOTP.requirements.AuthToken,
	})
}

func SendOtp(phone string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo("+91" + phone)
	params.SetChannel("sms")
	res, err := tw.VerifyV2.CreateVerification(twilioOTP.requirements.ServiceSid, params)
	if err != nil {
		return "", err
	}
	return *res.Sid, nil
}

func VerifyOtp(phone string, otp string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(otp)
	res, err := tw.VerifyV2.CreateVerificationCheck(twilioOTP.requirements.ServiceSid, params)
	if err != nil {
		return err
	}

	if *res.Status == "approved" {
		return nil
	}
	return errors.New("failed to verify otp")
}
