package createtransaction

import (
	"context"
	"testing"

	"github.com/flaviomdutra/ms-wallet-core/internal/entity"
	"github.com/flaviomdutra/ms-wallet-core/internal/event"
	"github.com/flaviomdutra/ms-wallet-core/internal/usecase/mocks"
	"github.com/flaviomdutra/ms-wallet-core/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCaseExecute(t *testing.T) {
	clientFrom, _ := entity.NewClient("Jack Reacher", "jack.reacher@email.com")
	clientTo, _ := entity.NewClient("John Wick", "jonh.wick@email.com")

	accountFrom := entity.NewAccount(clientFrom)
	accountTo := entity.NewAccount(clientTo)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	mockTransaction := &mocks.TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIDFrom: accountFrom.ID,
		AccountIDTo:   accountTo.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreated()
	ctx := context.Background()

	uc := NewCreateTransactionUseCase(mockUow, dispatcher, event)
	output, err := uc.Execute(ctx, input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
