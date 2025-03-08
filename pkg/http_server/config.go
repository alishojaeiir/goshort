package httpserver

import "time"

type Config struct {
	Port               int           `koanf:"port"`
	Cors               Cors          `koanf:"cors"`
	ShutDownCtxTimeout time.Duration `koanf:"shutdown_context_timeout"`
}

type Cors struct {
	AllowOrigins []string `koanf:"allow_origins"`
}
