package createtransaction

import (
	"testing"

	"github.com/flaviomdutra/ms-wallet-core/internal/entity"
	"github.com/flaviomdutra/ms-wallet-core/internal/event"
	"github.com/flaviomdutra/ms-wallet-core/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUseCaseExecute(t *testing.T) {
	clientFrom, _ := entity.NewClient("Jack Reacher", "jack.reacher@email.com")
	clientTo, _ := entity.NewClient("John Wick", "jonh.wick@email.com")

	accountFrom := entity.NewAccount(clientFrom)
	accountTo := entity.NewAccount(clientTo)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("FindByID", accountFrom.ID).Return(accountFrom, nil)
	mockAccount.On("FindByID", accountTo.ID).Return(accountTo, nil)

	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIDFrom: accountFrom.ID,
		AccountIDTo:   accountTo.ID,
		Amount:        100,
	}

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreated()
	uc := NewCreateTransactionUseCase(mockTransaction, mockAccount, dispatcher, event)
	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)

	mockAccount.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)

	mockAccount.AssertNumberOfCalls(t, "FindByID", 2)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)
}
