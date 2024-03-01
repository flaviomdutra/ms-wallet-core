package createclient

import (
	"testing"

	"github.com/flaviomdutra/ms-wallet-core/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClientUseCaseExecute(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	uc := NewCreateClientUseCase(m)

	input := CreateClientInputDTO{
		Name:  "Jack Reacher",
		Email: "jack.reacher@email.com"}

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Email, output.Email)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
