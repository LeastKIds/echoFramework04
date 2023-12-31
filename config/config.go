package config

import "fmt"

type EnvConfig struct {
	AppName string `envconfig:"APP_NAME" required:"true"`
	AppHost string `envconfig:"APP_HOST" required:"true"`
	AppPort string `envconfig:"APP_PORT" required:"true"`

	DBDriver  string `envconfig:"DB_DRIVER" required:"true"`
	DBHost    string `envconfig:"DB_HOST" required:"true"`
	DBPort    string `envconfig:"DB_PORT" required:"true"`
	DBSSLMode string `envconfig:"DB_SSLMODE" required:"true"`
	DBName    string `envconfig:"DB_NAME" requried:"true"`
	DBUser    string `envconfig:"DB_USER" requried:"true"`
	DBPass    string `envconfig:"DB_PASS" requried:"true"`
}

func (c *EnvConfig) DataSourceName() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", c.DBDriver, c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName, c.DBSSLMode)
}
