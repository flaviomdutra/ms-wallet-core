package createaccount

import (
	"testing"

	"github.com/flaviomdutra/ms-wallet-core/internal/entity"
	"github.com/flaviomdutra/ms-wallet-core/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCaseExecute(t *testing.T) {
	client, _ := entity.NewClient("Jack Reacher", "jack.reacher@email.com")
	clientMock := &mocks.ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMock, clientMock)
	input := CreateAccountInputDTO{ClientID: client.ID}

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)

	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)

	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
