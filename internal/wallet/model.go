package wallet

type WithdrawRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
	Source string  `json:"source"`
}

type AddBalanceRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
	Source string  `json:"source"`
}

type BalanceResponse struct {
	Balance           float64 `json:"balance"`
	LastTransactionID int64   `json:"last_transaction_id,omitempty"`
	UpdatedAt         string  `json:"updated_at,omitempty"`
	TransactionID     int64   `json:"transaction_id,omitempty"`
}