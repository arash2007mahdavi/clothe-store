package configs

import (
	"os"
	"github.com/spf13/viper"
)

func GetConfig() *Config {
	path := GetPath(os.Getenv("APP_ENV"))
	viper := LoadConfig(path, "yml")
	cfg := ParseConfig(viper)
	return cfg
}

func GetPath(env string) string {
	if env == "docker" {
		return "../configs/docker"
	} else if env == "production" {
		return "../configs/production"
	} else {
		return "../configs/development"
	}
}
func LoadConfig(filename string, filetype string) *viper.Viper {
	viper := viper.New()
	viper.SetConfigName(filename)
	viper.SetConfigType(filetype)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic("Problrm In Reading Configs")
	}
	return viper
}

func ParseConfig(viper *viper.Viper) *Config {
	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic("Problem In Parsing Configs")
	}
	return &cfg
}