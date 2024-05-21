package runner

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	config2 "github.com/yusologia/go-core/config"
	"github.com/yusologia/go-core/router"
	"reflect"
	"runtime"
	router2 "service/router"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "route-list",
		Long: "Running Route List",
		Run: func(cmd *cobra.Command, args []string) {
			config2.InitEnv()

			newRouter := mux.NewRouter()
			router.RegisterRouter(newRouter, router2.Register)

			newRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
				path, err := route.GetPathTemplate()

				if err != nil {
					return err
				}
				methods, _ := route.GetMethods()
				handlerName := getFunctionRouterName(route.GetHandler())

				if handlerName != "Unknown" {
					fmt.Println(methods, " || ", path, " || ", handlerName)

				}
				return nil
			})
		},
	})
}

func getFunctionRouterName(f interface{}) string {
	if f == nil {
		return "Unknown"
	}

	ptr := reflect.ValueOf(f).Pointer()
	funcName := runtime.FuncForPC(ptr).Name()

	if len(funcName) >= 4 && funcName[len(funcName)-3:] == "-fm" {
		funcName = funcName[:len(funcName)-3]
	}

	return funcName
}
