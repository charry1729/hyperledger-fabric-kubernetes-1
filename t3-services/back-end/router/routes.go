package router

import (
	"github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/models"
	HomeHandler "github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/routes/home"
	StatusHandler "github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/routes/status"
)

func  GetRoutes() models.Routes{
	return models.Routes{
		models.Route{
			Name: "Home",
			Method: "GET",
			Pattern: "/",
			HandlerFunc: HomeHandler.Index
		},
		models.Route{
			Name: "Status",
			Method: "GET",
			Pattern: "/status",
			HandlerFunc: StatusHandler.Index
		}
	}

}