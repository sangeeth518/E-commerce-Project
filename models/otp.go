package models

// request for otp verification
type OTPData struct {
	Number string `json:"number,omitempty"`
}
type VerifyData struct {
	PhoneNumber string `json:"phone,omitempty" validate:"required"`
	Code        string `json:"code,omitempty" validate:"required"`
}
