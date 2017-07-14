package steamtrade

import (
	"net/http"
	"log"
	"time"
	"context"
	"github.com/gorilla/mux"
)


var Server struct {
	Server *http.Server
	Handler *mux.Router
}


func GetRouters() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", MethodHelp).Methods("GET")
	r.HandleFunc("/stop", MethodStop).Methods("POST")
	r.HandleFunc("/offer/list", MethodGetOfferList).Methods("GET")
	r.HandleFunc("/offer/{id}", MethodGetOffer).Methods("GET")
	r.HandleFunc("/offer/{id}", MethodCancelOffer).Methods("DELETE")
	return r
}


func InitializeServer() {
	if Server.Handler == nil {
		Server.Handler = GetRouters()
	}
	Server.Server = &http.Server{
		Addr: ":" + Config.ServerPort,
		Handler: Server.Handler,
	}
}


func StartServer() error {
	InitializeServer()
	log.Fatal(Server.Server.ListenAndServe())
	return nil
}


func StopServer() {
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	Server.Server.Shutdown(ctx)
}
