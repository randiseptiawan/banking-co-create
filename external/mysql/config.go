package mysql

import "github.com/rysmaadit/go-template/config"

type ClientConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}

func Init_db() *ClientConfig {
	init := config.Init()
	config := ClientConfig{
		Username: init.DBUsername,
		Password: init.DBPassword,
		Host:     init.DBHost,
		Port:     init.DBPort,
		DBName:   init.DBName,
	}
	return &config
}
