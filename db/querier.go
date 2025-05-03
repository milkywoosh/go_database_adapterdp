package db

import (
	"context"
)

type CreateUserTxResult struct {
	Users
}

type CreateUserParams struct {
	Username, Email, Firstname, Lastname, Password string
}

type Querier interface {
	UserDo
	PurchaseDo

	// AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	// CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	// CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	// CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	// CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	// CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	// DeleteAccount(ctx context.Context, id int64) error
	// GetAccount(ctx context.Context, id int64) (Account, error)
	// GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	// GetEntry(ctx context.Context, id int64) (Entry, error)
	// GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	// GetTransfer(ctx context.Context, id int64) (Transfer, error)
	// GetUser(ctx context.Context, username string) (User, error)
	// ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	// ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	// ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	// UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	// UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	// UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
}

type UserDo interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
}

type PurchaseDo interface {
	CreatePurchaseHistory(ctx context.Context, arg CreatePurchaseHistoryParams) (PurchaseHistory, error)
	AddListBook(ctx context.Context, arg CreateBookToPurchaseParams) (BookToPurchase, error)
	DeletePurchase(ctx context.Context, arg ...interface{}) error
	FinalizePurchase() error
	AdjustStockBook(ctx context.Context, bookID int, corrector int) error
}

var _ Querier = (*Queries)(nil)
