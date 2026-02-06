package router

import (
	"net/http"

	"go-wallet/internal/config"
	"go-wallet/internal/auth"
	"go-wallet/internal/wallet"
)

func Register(
	cfg config.Config,
	walletHandler *wallet.Handler,
) http.Handler {

	mux := http.NewServeMux()

	wallet.RegisterRoutes(mux, walletHandler)

	keyMap := make(map[string]struct{})
	for _, k := range cfg.APIKeys {
		if k != "" {
			keyMap[k] = struct{}{}
		}
	}

	return middleware.APIKeyAuth(keyMap)(mux)
}