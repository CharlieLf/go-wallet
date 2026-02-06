package wallet

import (
	"context"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) Withdraw(
	ctx context.Context,
	userID string,
	amount float64,
	source string,
) (*BalanceResponse, error) {

	row := r.DB.QueryRowContext(
		ctx,
		"CALL WithdrawUserBalance(?,?,?)",
		userID,
		amount,
		source,
	)

	var resp BalanceResponse

	err := row.Scan(
		&resp.Balance,
		&resp.TransactionID,
	)

	return &resp, err
}

func (r *Repository) AddBalance(
	ctx context.Context,
	userID string,
	amount float64,
	source string,
) (*BalanceResponse, error) {

	row := r.DB.QueryRowContext(
		ctx,
		"CALL AddUserBalance(?,?,?)",
		userID,
		amount,
		source,
	)

	var resp BalanceResponse

	err := row.Scan(
		&resp.Balance,
		&resp.TransactionID,
	)

	return &resp, err
}

func (r *Repository) GetBalance(
	ctx context.Context,
	userID string,
) (*BalanceResponse, error) {

	row := r.DB.QueryRowContext(
		ctx,
		"CALL GetUserBalance(?)",
		userID,
	)

	var resp BalanceResponse

	err := row.Scan(
		&resp.Balance,
		&resp.UpdatedAt,
		&resp.LastTransactionID,
	)

	return &resp, err
}