package config

import "fmt"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Host string
	Port int
	User string
	Name string
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host: "0.0.0.0",
			Port: 3306,
			User: "root",
			Name: "todoapp",
		},
	}
}
