package config

type EnvConfig struct {
	AppName string `envconfig:"APP_NAME" required:"true"`
	AppHost string `envconfig:"APP_HOST" required:"true"`
	AppPort string `envconfig:"APP_PORT" required:"true"`
}
