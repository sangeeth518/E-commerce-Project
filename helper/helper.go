package helper

import (
	"errors"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient

func TwilioSetup(username string, password string) {
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: username,
		Password: password,
	})
}

func TwilioSndOtp(phone string, serviceId string) (string, error) {
	to := "+91" + phone
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(serviceId, params)
	if err != nil {
		return "", err
	}
	return *resp.Sid, nil

}

func TwilioVerifyOtp(phone string, otp string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(otp)
	resp, err := client.VerifyV2.CreateVerificationCheck(os.Getenv("SERVICE_SID"), params)
	if err != nil {
		return err
	}
	if *resp.Status == "approved" {
		return nil

	}
	return errors.New("failed to validate otp")
}
