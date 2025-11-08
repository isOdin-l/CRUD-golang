package configs

type Config struct {
	SERVER_PORT     string `env:"SERVER_PORT"`
	DB_PASSWORD     string `env:"DB_PASSWORD"`
	DB_USERNAME     string `env:"DB_USERNAME"`
	DB_HOST         string `env:"DB_HOST"`
	DB_PORT         string `env:"DB_PORT"`
	DB_NAME         string `env:"DB_NAME"`
	SALT            string `env:"SALT"`
	JWT_SIGNING_KEY string `env:"JWT_SIGNING_KEY"`
}

type InternalConfig struct {
	SALT            string `env:"SALT"`
	JWT_SIGNING_KEY string `env:"JWT_SIGNING_KEY"`
}
