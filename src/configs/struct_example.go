package configs

type Config struct {
	Server  ServerConfig
	Store   StoreConfig
	Logging LoggingConfig
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
