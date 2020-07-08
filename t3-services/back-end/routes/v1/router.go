package v1

import (
	"log"
	"net/http"
	"github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/models"
	UsersHandler "github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/models/v1/users"
   )

func Middleware() func(http.Handler) http.Handler{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
			log.Println("Inside v1 middleware")
			next.ServeHTTP(w,r)
		})
	}
}

func GetRoutes() map[string] models.SubRoutePackage{
	return map[string] models.SubRoutePackage{
		"/v1": {
			Middleware: Middleware(),
			Routes:  models.Routes{
				models.Route{Name: "UsersIndex", Method: "GET", Pattern: "/users",HandlerFunc : UsersHandler.Index()},
				models.Route{Name: "UsersStore", Method: "POST", Pattern: "/users",HandlerFunc : UsersHandler.Store()},
				models.Route{Name: "UsersUpdate", Method: "PUT", Pattern: "/users/{id}",HandlerFunc : UsersHandler.Update()},
				models.Route{Name: "UsersUpdate", Method: "PATCH", Pattern: "/users/{id}",HandlerFunc : UsersHandler.Update()},
				models.Route{Name: "UsersDestroy", Method: "DELETE", Pattern: "/users/{id}",HandlerFunc : UsersHandler.Destroy()},

			},
		},
	}
}