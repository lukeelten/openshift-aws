package aws

type AwsConfig struct {
	Region string
	KeyId string
	SecretKey string
}


func NewConfig(region string, keyid string, secretkey string) *AwsConfig {
	config := AwsConfig{
		Region: region,
		KeyId: keyid,
		SecretKey: secretkey,
	}

	return &config
}