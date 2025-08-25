package config

type Config struct {
	Release        string `yaml:"release" env:"RELEASE" env-default:"local"`
	AppEnvironment string `yaml:"app_environment" env:"APP_ENVIRONMENT" env-default:"local"`
	HTTP           HTTP   `yaml:"http"`
}

type HTTP struct {
	Port           string `yaml:"port" env:"HTTP_SERVER_PORT" env-default:"42069"`
	ReadTimeoutMS  int64  `yaml:"read_timeout_ms" env:"HTTP_SERVER_READ_TIMEOUT_MS" env-default:"10000"`
	WriteTimeoutMS int64  `yaml:"write_timeout_ms" env:"HTTP_SERVER_WRITE_TIMEOUT_MS" env-default:"10000"`
}
