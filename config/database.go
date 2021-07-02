package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Host    string
	Port    int
	User    string
	Name    string
	Charset string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host:    "0.0.0.0",
			Port:    3306,
			User:    "root",
			Name:    "todoapp",
			Charset: "utf8",
		},
	}
}
