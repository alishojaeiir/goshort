package config

import (
	"github.com/alishojaeiir/GoShort/pkg/database"
	httpserver "github.com/alishojaeiir/GoShort/pkg/http_server"
	"github.com/alishojaeiir/GoShort/pkg/logger"
	"time"
)

type Config struct {
	Database             database.Config   `koanf:"db"`
	Logger               logger.Config     `koanf:"logger"`
	Server               httpserver.Config `koanf:"server"`
	TotalShutdownTimeout time.Duration     `koanf:"total_shutdown_timeout"`
}
