package server 

import (
	"strconv"
	"net/http"
	"time"
	"log"
	"os"
	"github.com/gorilla/handlers"
	"github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/router"
)

type Server struct{
	Initialized bool 
	Addr string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	MaxHeaderBytes int
	Port int
	HTTPServer *http.Server
}

func (s *Server) Init(port int) error  {
	s.Addr = ":" + strconv.Itoa(port)
	s.Port = port

	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second

	s.MaxHeaderBytes = 1 << 20
	r := router.GetRouter()

	handler := handlers.LoggingHandler(os.Stdout,handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET","PUT","PATCH","POST","DELETE"}),
		handlers.AllowedMethods([]string{"Content-Type","Origin","Cache-Control"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
		
	)(r.GetRawRouter()))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)
	s.HTTPServer = &http.Server{
		Addr : s.Addr,
		Handler: handler,
		ReadHeaderTimeout: s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
		MaxHeaderBytes: s.MaxHeaderBytes,

	}

	s.Initialized = true
	return nil
}


func (s *Server) Start(){
	if !s.Initialized{
		panic("A server needs to be initialized before its started")
	}
	log.Printf("Server started on port: %d", s.Port)
	log.Fatal(s.HTTPServer.ListenAndServe())

}