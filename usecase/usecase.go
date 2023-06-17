package usecase

import (
	"context"
	"errors"
	"log"
	"transactionlabo/domain"
	"transactionlabo/port"
)

type Usecase struct {
	DBRepository   port.DBRepository
	TxDBRepository port.TxRepository
}

func NewUsecase(dbRepository port.DBRepository, txDBRepository port.TxRepository) *Usecase {
	return &Usecase{
		DBRepository:   dbRepository,
		TxDBRepository: txDBRepository,
	}
}

func (r *Usecase) CreateUser(ctx context.Context) error {
	return r.TxDBRepository.Transaction(ctx, func(ctx context.Context, db port.DBRepository) error {
		user := &domain.User{
			Name: "test_user",
		}
		userRes, err := db.CreateUser(ctx, user)
		if err != nil {
			return err
		}
		role := &domain.Role{
			Name:   "test_role",
			UserID: userRes.ID,
		}
		roleRes, err := db.CreateRole(ctx, role)
		if err != nil {
			return err
		}
		log.Println("userRes", userRes, "roleRes", roleRes)
		return errors.New("rollback")
	})
}
