package main

import (
	"database/sql"
	"fmt"

	"github.com/flaviomdutra/ms-wallet-core/internal/database"
	"github.com/flaviomdutra/ms-wallet-core/internal/event"
	createaccount "github.com/flaviomdutra/ms-wallet-core/internal/usecase/create_account"
	createclient "github.com/flaviomdutra/ms-wallet-core/internal/usecase/create_client"
	createtransaction "github.com/flaviomdutra/ms-wallet-core/internal/usecase/create_transaction"
	"github.com/flaviomdutra/ms-wallet-core/internal/web"
	"github.com/flaviomdutra/ms-wallet-core/internal/web/webserver"
	"github.com/flaviomdutra/ms-wallet-core/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	// eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webserver.Start()
}