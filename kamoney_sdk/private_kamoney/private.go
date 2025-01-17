package private_kamoney

type PrivateRequestsInterface interface {
	AccountInfo() (struct{}, struct{})
}

type privateRequests struct {
	Email     string
	Password  string
	PublicKey string
	SecretKey string
	r         *RequestHandler
}

func NewPrivateRequests(email, password, publicKey, secretKey string, r *RequestHandler) PrivateRequestsInterface {
	return &privateRequests{
		Email:     email,
		Password:  password,
		PublicKey: publicKey,
		SecretKey: secretKey,
		r:         r,
	}
}
