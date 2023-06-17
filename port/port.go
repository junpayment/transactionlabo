package port

import (
	"context"
	"transactionlabo/domain"
)

type DBRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	CreateRole(ctx context.Context, role *domain.Role) (*domain.Role, error)
}

type TxRepository interface {
	Transaction(ctx context.Context, fn func(ctx context.Context, db DBRepository) error) error
}
