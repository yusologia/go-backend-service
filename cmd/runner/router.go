package runner

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/yusologia/go-core/v2"
	"github.com/yusologia/go-core/v2/pkg"
	"reflect"
	"runtime"
	"service/internal/app/api"
	"service/internal/app/privateapi"
	"strconv"
	"strings"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "logia:router",
		Long: "Running Route List",
		Run: func(cmd *cobra.Command, args []string) {
			logiapkg.InitDevMode()

			newRoute := mux.NewRouter()
			logiacore.RegisterRouter(newRoute, api.Register, privateapi.Register)

			methodLen := 0
			pathLen := 0

			routeLists := make([][]string, 0)
			newRoute.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
				path, err := route.GetPathTemplate()
				if err != nil {
					return err
				}

				methods, _ := route.GetMethods()
				handlerName := getFunctionName(route.GetHandler())

				if handlerName != nil {
					method := fmt.Sprintf("[%s]", strings.Join(methods, ", "))
					if methodLenStr := len(method); methodLenStr > methodLen {
						methodLen = methodLenStr
					}

					if pathLenStr := len(path); pathLenStr > pathLen {
						pathLen = pathLenStr
					}

					routeLists = append(routeLists, []string{method, path, *handlerName})
				}

				return nil
			})

			methodFormat := "%-" + strconv.Itoa(methodLen+3) + "s"
			pathFormat := "%-" + strconv.Itoa(pathLen+5) + "s"

			for _, routeList := range routeLists {
				printMethod := fmt.Sprintf(methodFormat, routeList[0])
				printPath := fmt.Sprintf(pathFormat, routeList[1])

				fmt.Printf("%s %s %s\n", printMethod, printPath, routeList[2])
				logiapkg.LogDebug(routeList)
			}
		},
	})
}

func getFunctionName(f interface{}) *string {
	if f == nil {
		return nil
	}

	ptr := reflect.ValueOf(f).Pointer()
	funcName := runtime.FuncForPC(ptr).Name()

	if len(funcName) >= 4 && funcName[len(funcName)-3:] == "-fm" {
		funcName = funcName[:len(funcName)-3]
	}

	return &funcName
}
