package database

type Config struct {
	Driver           string `koanf:"driver"`
	Host             string `koanf:"host"`
	Port             int    `koanf:"port"`
	Username         string `koanf:"username"`
	Password         string `koanf:"password"`
	DBName           string `koanf:"db_name"`
	SSLMode          string `koanf:"ssl_mode"`
	MaxIdleConns     int    `koanf:"max_idle_conns"`
	MaxOpenConns     int    `koanf:"max_open_conns"`
	ConnMaxLifetime  int    `koanf:"conn_max_lifetime"`
	PathOfMigrations string `koanf:"path_of_migrations"`
}
