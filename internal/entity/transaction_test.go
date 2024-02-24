package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	clientFrom, _ := NewClient("John Doe", "jonh.doe@email.com")
	clientTo, _ := NewClient("Jane Doe", "jane.doe@email.com")
	accountFrom := NewAccount(clientFrom)
	accountTo := NewAccount(clientTo)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)

	transaction, err := NewTransaction(accountFrom, accountTo, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)

	assert.Equal(t, 900.0, accountFrom.Balance)
	assert.Equal(t, 1100.0, accountTo.Balance)
}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	clientFrom, _ := NewClient("John Doe", "jonh.doe@email.com")
	clientTo, _ := NewClient("Jane Doe", "jane.doe@email.com")
	accountFrom := NewAccount(clientFrom)
	accountTo := NewAccount(clientTo)

	accountFrom.Credit(1000)
	accountTo.Credit(1000)

	transaction, err := NewTransaction(accountFrom, accountTo, 2000)

	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")

	assert.Nil(t, transaction)

	assert.Equal(t, 1000.0, accountFrom.Balance)
	assert.Equal(t, 1000.0, accountTo.Balance)
}
