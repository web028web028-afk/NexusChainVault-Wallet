package main

import (
	"encoding/json"
	"net/http"
)

type APIRouter struct {
	wallet *NexusWallet
	mux    *http.ServeMux
}

func NewAPIRouter(w *NexusWallet) *APIRouter {
	r := &APIRouter{wallet: w, mux: http.NewServeMux()}
	r.registerRoutes()
	return r
}

func (a *APIRouter) registerRoutes() {
	a.mux.HandleFunc("/api/wallet/info", a.handleWalletInfo)
	a.mux.HandleFunc("/api/tx/send", a.handleSendTransaction)
}

func (a *APIRouter) handleWalletInfo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(a.wallet)
}

func (a *APIRouter) handleSendTransaction(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"pending"}`))
}

func (a *APIRouter) StartServer(addr string) error {
	return http.ListenAndServe(addr, a.mux)
}
