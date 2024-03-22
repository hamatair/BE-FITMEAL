package service

import (
	"errors"

	"intern-bcc/entity"
	"intern-bcc/internal/repository"
	"intern-bcc/model"
	"intern-bcc/pkg/midtranss"

	"github.com/google/uuid"
)

type TopUpServiceI interface {
	ConfirmedTopUp(id string, data map[string]interface{}) error
	InitializeTopUp(req model.TopUpReq) (model.TopUpRes, error)
}

type TopUpservices struct {
	// notificationService notificationService
	TopUpRepository repository.TopUpRepositoryI
	midtransService midtranss.MidtransServiceI
	Account         repository.UserRepositoryInterface
}

func NewTopUp( /*notificationService notificationService,*/ topUpRepository repository.TopUpRepositoryI, midtransService midtranss.MidtransServiceI, user *repository.Repository) TopUpServiceI {
	return &TopUpservices{
		// notificationService: notificationService,
		TopUpRepository: topUpRepository,
		midtransService: midtransService,
		Account:         user.UserRepository,
	}
}

// InitializeTopUp implements TopUpServiceI.
func (t *TopUpservices) InitializeTopUp(req model.TopUpReq) (model.TopUpRes, error) {
	topUp := entity.TopUp{
		ID:     uuid.New(),
		UserID: req.UserId,
		Status: 0,
		Amount: uint(req.Amount),
	}

	err := t.midtransService.GenerateSnapUrl(&topUp)
	if err != nil {
		return model.TopUpRes{}, err
	}

	err = t.TopUpRepository.Insert(&topUp)
	if err != nil {
		return model.TopUpRes{}, err
	}

	return model.TopUpRes{
		SnapUrl: topUp.SnapUrl,
	}, nil
}

// ConfirmedTopUp implements TopUpServiceI.
func (t *TopUpservices) ConfirmedTopUp(id string, data map[string]interface{}) error {
	topUp, err := t.TopUpRepository.FindById(id)
	if err != nil {
		return err
	}

	if topUp == (entity.TopUp{}) {
		return errors.New("topUp request Not found")
	}

	topUp.Status = 1

	err = t.TopUpRepository.Update(&topUp)
	if err != nil {
		return err
	}

	err = t.midtransService.VerifyPayment(data)
	if err != nil {
		return err
	}

	account, err := t.Account.GetUser(model.UserParam{
		ID: topUp.ID,
	})
	if err != nil {
		return err
	}

	account.Balance += topUp.Amount

	err = t.Account.UpdateUser(account, model.UserParam{
		ID: topUp.ID,
	})
	if err != nil {
		return err
	}

	// data := map[string]string{
	// 	"amount": fmt.Sprintf("%2.f", topUp.Amount),
	// }
	return nil
}

