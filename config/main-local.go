package config

import (
	httpserver "github.com/alishojaeiir/GoShort/pkg/http_server"
	"time"
)

type Config struct {
	TotalShutdownTimeout time.Duration     `koanf:"total_shutdown_timeout"`
	Server               httpserver.Config `koanf:"server"`
}
