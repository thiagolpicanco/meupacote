package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Username string
	Password string
	Host     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Username: "admin",
			Password: "admin",
			Host:     "ds225028.mlab.com:25028/meupacote",
		},
	}
}
