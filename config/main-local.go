package config

import (
	httpserver "github.com/alishojaeiir/GoShort/pkg/http_server"
	"github.com/alishojaeiir/GoShort/pkg/logger"
	"time"
)

type Config struct {
	TotalShutdownTimeout time.Duration     `koanf:"total_shutdown_timeout"`
	Server               httpserver.Config `koanf:"server"`
	Logger               logger.Config     `koanf:"logger"`
}
