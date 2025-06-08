package config

import (
	"github.com/rs/cors"
	"github.com/yusologia/go-core/v2/pkg"
)

var (
	CorsOptions cors.Options
)

func InitCors() {
	CorsOptions.AllowedOrigins = []string{logiapkg.HostFull}
	CorsOptions.AllowCredentials = false
	CorsOptions.AllowedMethods = []string{"GET", "POST", "PUT", "DELETE"}
	CorsOptions.AllowedHeaders = []string{"*"}
}
