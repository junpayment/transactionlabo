package repository

import (
	"context"
	"transactionlabo/domain"
	"transactionlabo/port"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(dsn string) (*DBRepository, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DBRepository{
		db: db,
	}, nil
}

func NewTxDBRepository(db *gorm.DB) *DBRepository {
	return &DBRepository{
		db: db,
	}
}

func (r *DBRepository) DB() *gorm.DB {
	return r.db
}

func (r *DBRepository) Transaction(
	ctx context.Context,
	fn func(ctx2 context.Context, db port.DBRepository) error,
) error {
	return r.db.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			db := NewTxDBRepository(tx)
			return fn(ctx, db)
		},
	)
}

func (r *DBRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *DBRepository) CreateRole(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	if err := r.db.WithContext(ctx).Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}
