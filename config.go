package main

type Config struct {
	Github          Github
	ImpersonateTags bool `flag:"impersonatetags" env:"IMPERSONATE_TAGS"`
	Port            int  `flag:"port" env:"HTTP_PORT"`
}

type Github struct {
	AppId            int64  `flag:"appid" env:"GITHUB_APP_ID"`
	WebhookSecretKey string `flag:"webhooksecret" env:"GITHUB_WEBHOOK_SECRET"`
	PrivateKey       string `flag:"privatekey" env:"GITHUB_PRIVATE_KEY"`
	EnterpriseURL    string `flag:"enterpriseurl" env:"GITHUB_ENTERPRISE_URL"`
}

var (
	cfg = &Config{
		Github:          Github{},
		ImpersonateTags: false,
		Port:            8080,
	}
)
