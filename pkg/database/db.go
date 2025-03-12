package database

import (
	"database/sql"
	"fmt"
	"time"
)

type DSNBuilder interface {
	BuildDSN(config Config) string
}

type Database struct {
	DB      *sql.DB
	Dialect string
}

func Connect(config Config) (*Database, error) {

	dsnB, DSNErr := GetDSNBuilder(config.Driver)
	if DSNErr != nil {
		return nil, DSNErr
	}
	dsn := dsnB.BuildDSN(config)
	conn, err := sql.Open(config.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	version, vErr := getVersion(conn)
	if vErr != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(config.MaxIdleConns)
	conn.SetMaxOpenConns(config.MaxOpenConns)
	conn.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)

	fmt.Println("Database connection established successfully")
	fmt.Printf("The version of database is: %s\n", version)

	return &Database{DB: conn, Dialect: config.Driver}, err
}

func Close(conn *sql.DB) error {
	return conn.Close()
}

func getVersion(db *sql.DB) (string, error) {
	var res string
	err := db.QueryRow("SELECT version()").Scan(&res)
	if err != nil {
		return "", fmt.Errorf("error executing query: %v", err)
	}
	return res, nil
}
