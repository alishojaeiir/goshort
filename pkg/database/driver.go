package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // MySQL
	_ "github.com/lib/pq"              // PostgreSQL
	_ "github.com/mattn/go-sqlite3"    // SQLite
)

var DSNRegistry = make(map[string]DSNBuilder)

func RegisterDSNBuilder(dsn string, builder DSNBuilder) {
	DSNRegistry[dsn] = builder
}

func GetDSNBuilder(dsn string) (DSNBuilder, error) {
	builder, exists := DSNRegistry[dsn]
	if !exists {
		return nil, fmt.Errorf("database driver %s is not registered", dsn)
	}
	return builder, nil
}

func init() {
	RegisterDSNBuilder("postgres", &PostgresDSNBuilder{})
	RegisterDSNBuilder("mysql", &MySQLDSNBuilder{})
	RegisterDSNBuilder("sqlite", &SQLiteDSNBuilder{})
}

type PostgresDSNBuilder struct{}

func (p *PostgresDSNBuilder) BuildDSN(config Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode)
}

type MySQLDSNBuilder struct{}

func (m *MySQLDSNBuilder) BuildDSN(config Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=%s",
		config.Username, config.Password, config.Host, config.Port, config.DBName, config.SSLMode)
}

type SQLiteDSNBuilder struct{}

func (s *SQLiteDSNBuilder) BuildDSN(config Config) string {
	return config.DBName
}
