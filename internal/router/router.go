package router

import (
	"net/http"

	"go-wallet/internal/wallet"
)

func Register(walletHandler *wallet.Handler) http.Handler {

	mux := http.NewServeMux()

	// Register all API domain
	wallet.RegisterRoutes(mux, walletHandler)

	return mux
}