package models

// request for otp verification
type OTPData struct {
	Number string `json:"number,omitempty"`
}
