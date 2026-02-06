package wallet

import "context"

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Withdraw(ctx context.Context, req WithdrawRequest) (*BalanceResponse, error) {
	return s.repo.Withdraw(ctx, req.UserID, req.Amount, req.Source)
}

func (s *Service) AddBalance(ctx context.Context, req AddBalanceRequest) (*BalanceResponse, error) {
	return s.repo.AddBalance(ctx, req.UserID, req.Amount, req.Source)
}

func (s *Service) GetBalance(ctx context.Context, userID string) (*BalanceResponse, error) {
	return s.repo.GetBalance(ctx, userID)
}