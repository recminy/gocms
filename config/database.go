package config

import (
	"cms/pkg/config"
)

func init()  {
	config.Add("database", config.Map{
		"mysql": map[string]interface{}{
			"host": config.Env("DB_HOST", "127.0.0.1"),
			"port": config.Env("DB_PORT", 3306),
			"database": config.Env("DB_NAME", ""),
			"username": config.Env("DB_USER", ""),
			"password": config.Env("DB_PASSWORD", ""),
			"charset": "utf8mb4",

			"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 100),
			"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 25),
			"max_life_seconds": config.Env("DB_MAX_LIFE_SECONDS", 5*60),
		},
	})
}
