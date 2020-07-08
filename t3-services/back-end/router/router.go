package router

import ("github.com/gorilla/mux"
		"github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/models"
		V1Router "github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/routes/v1"
)

type Router struct {
	RawRouter *mux.Router
}

func (r Router) GetRawRouter() *mux.Router  {
	return r.RawRouter
	
}


function (r *Router) AttachSubRouterWithMiddleware(path string, subroutes models.Routes,middleware mux.MiddlewareFunc){
	SubRouter := r.RawRouter.PathPrefix(path).Subrouter()
	SubRouter.Use(middleware)
	for _, route := range subroutes{
		SubRouter.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(route.HandlerFunc)
	}
	for name,pack := range V1Router.GetRoutes(){
		r.AttachSubRouterWithMiddleware(name,pack.Routes,pack.Middleware)
	}

	return SubRouter



}
