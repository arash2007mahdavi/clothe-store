package configs

type Config struct {
	Server ServerConfig
	Store StoreConfig
}

type ServerConfig struct {
	Port int
}

type StoreConfig struct {
	Hat Product
	Shoes Product
	Pant Product
	Shirt Product
}

type Product struct {
	Amount int
	Price float64
	Currency string
}