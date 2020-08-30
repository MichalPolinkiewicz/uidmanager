package main

import (
	"net/http"
	config "uidmanager/config"
	"uidmanager/handlers"
)


func main() {
	cfg := config.NewConfiguration()

	mux := http.NewServeMux()
	mux.HandleFunc("/setuid", handlers.SetBidderUIDEndpoint(cfg))
	if err := http.ListenAndServe(cfg.Port, mux); err != nil{

	}
}