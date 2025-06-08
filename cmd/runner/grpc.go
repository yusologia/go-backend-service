package runner

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/v2/grpc"
	"github.com/yusologia/go-core/v2/pkg"
	"google.golang.org/grpc/reflection"
	"log"
	"os"
	"service/internal/app/grpc"
	"service/internal/pkg/config"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "logia:grpc",
		Long: "Running gRPC",
		Run: func(cmd *cobra.Command, args []string) {
			logiapkg.InitDevMode()

			config.InitTZ()

			DBClose := config.InitDB()
			defer DBClose()

			addr := fmt.Sprintf("%s", os.Getenv("GRPC_HOST"))

			server := logiagrpc.GRPCServer{}
			server.NewServer(addr)

			grpc.Register(&server)
			reflection.Register(server.Server)

			fmt.Println(fmt.Sprintf("gRPC server is running: %s", addr))
			if err := server.Serve(); err != nil {
				log.Fatalf("Failed to server: %v", err)
			}
		},
	})
}
