package aws_aegis_auth

type AuthAWS struct {
	AccountNumber string `json:"accountNumber"`
	Region        string `json:"region"`
	AccessKey     string `json:"accessKey"`
	SecretKey     string `json:"secretKey"`
}
