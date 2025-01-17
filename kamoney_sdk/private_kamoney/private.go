package private

import "github.com/kamoney/sdk_golang/kamoney_sdk/request"

type PrivateRequestsInterface interface {
	AccountInfo() (struct{}, struct{})
}

type privateRequests struct {
	Email     string
	Password  string
	PublicKey string
	SecretKey string
	r         *request.RequestHandler
}

func NewPrivateRequests(email, password, publicKey, secretKey string, r *request.RequestHandler) PrivateRequestsInterface {
	return &privateRequests{
		Email:     email,
		Password:  password,
		PublicKey: publicKey,
		SecretKey: secretKey,
		r:         r,
	}
}
