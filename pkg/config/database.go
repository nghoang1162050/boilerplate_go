package config

type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
	Timezone string `mapstructure:"DB_TIMEZONE"`
	Platform string `mapstructure:"DB_PLATFORM"`
	SslMode  string `mapstructure:"DB_SSL_MODE"`
}
