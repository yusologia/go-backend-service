package config

import (
	"github.com/rs/cors"
	"github.com/yusologia/go-core/config"
)

var (
	CorsOptions cors.Options
)

func InitCors() {
	CorsOptions.AllowedOrigins = []string{config.HostFull}
	CorsOptions.AllowCredentials = false
	CorsOptions.AllowedMethods = []string{"GET", "POST", "PUT", "DELETE"}
	CorsOptions.AllowedHeaders = []string{"*"}
}
