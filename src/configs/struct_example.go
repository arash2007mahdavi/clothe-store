package configs

import "time"

type Config struct {
	Server   ServerConfig
	Store    StoreConfig
	Logging  LoggingConfig
	Postgres PostgresConfig
	Otp      OtpConfig
	Jwt      JwtConfig
}

type JwtConfig struct {
	Secret                     string
	RefreshSecret              string
	AccessTokenExpireDuration  time.Duration
	RefreshTokenExpireDuration time.Duration
}

type OtpConfig struct {
	Expire time.Duration
	Digit  int
}

type PostgresConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	DbName          string
	SslMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type ServerConfig struct {
	Port int
}

type StoreConfig struct {
	Hat   Product
	Shoes Product
	Pant  Product
	Shirt Product
}

type LoggingConfig struct {
	Path     string
	LogLevel string
	Logger   string
}

type Product struct {
	Price    float64
	Currency string
}
