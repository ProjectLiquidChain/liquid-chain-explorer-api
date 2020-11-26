package api

import (
	"log"
	"net/http"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/server/surf"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/rs/cors"
)

// Server contains all info to serve an explorer api server
type Server struct {
	url        string
	rpcServer  *rpc.Server
	httpServer *http.Server
	Router     *mux.Router
}

// New return an new instance of Server
func New(url, dbURL, nodeURL, storagePath string) Server {
	api := Server{
		url: url,
	}
	api.setupServer()
	api.registerServices(dbURL, nodeURL, storagePath)
	api.setupRouter()
	return api
}

func (api *Server) setupServer() {
	server := rpc.NewServer()
	server.RegisterCodec(json2.NewCodec(), "application/json")
	api.rpcServer = server
}

func (api *Server) setupRouter() {
	api.Router = mux.NewRouter()
	api.Router.Handle("/", api.rpcServer).Methods("POST")
	api.httpServer = &http.Server{
		Handler: cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
			AllowedMethods:   []string{"POST", "DELETE", "PUT", "GET", "HEAD", "OPTIONS"},
		}).Handler(api.Router),
		Addr: api.url,
	}
}

func (api *Server) registerServices(dbURL, nodeURL, storagePath string) {
	if api.rpcServer == nil {
		panic("api.registerServices call without api.server")
	}
	if err := api.rpcServer.RegisterService(surf.New(dbURL, nodeURL, storagePath), "surf"); err != nil {
		panic(err)
	}
}

// Serve starts the server to serve request
func (api *Server) Serve() error {
	log.Println("Server is ready at", api.url)
	err := api.httpServer.ListenAndServe()
	return err
}

// Close will immediately stop the server without waiting for any active connection to complete
// For gracefully shutdown please implement another function and use Server.Shutdown()
func (api *Server) Close() {
	log.Println("Closing server")
	if api.httpServer != nil {
		err := api.httpServer.Close()
		if err != nil {
			panic(err)
		}
	}
}
