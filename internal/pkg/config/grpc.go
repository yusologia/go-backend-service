package config

import (
	"github.com/yusologia/go-core/v2/pkg"
	"os"
	"time"
)

var (
	TestingRPC string
)

func InitRPC() {
	logiapkg.RPCDialTimeout = 5 * time.Second
	TestingRPC = os.Getenv("GRPC_TESTING_HOST")
}
