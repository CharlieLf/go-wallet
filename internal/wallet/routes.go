package wallet

import "net/http"

func RegisterRoutes(mux *http.ServeMux, handler *Handler) {

	mux.HandleFunc("/api/v1/wallet/withdraw", handler.Withdraw)
	mux.HandleFunc("/api/v1/wallet/add", handler.AddBalance)
	mux.HandleFunc("/api/v1/wallet/balance", handler.GetBalance)

}