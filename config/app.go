package config

import (
	"cms/pkg/config"
)

func init()  {
	config.Add("app", config.Map{
		"port": config.Env("APP_PORT", "80"),
		"name": config.Env("APP_NAME", "myApp"),
		"env": config.Env("APP_ENV", "production"),
		"url": config.Env("APP_URL", "http://localhost"),
		"debug": config.Env("APP_DEBUG", false),
		"key": config.Env("APP_KEY", "daa60df2f8080505402d0ce37fbfcff0"),
	})
}