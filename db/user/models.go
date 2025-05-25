package db

type Book struct {
	BookID   int
	Title    string
	Price    float64
	StockQty int
	AuthorID int
}

type Users struct {
	Username,
	Email,
	Firstname,
	Lastname string
}

type CreateUserTxResult struct {
	Users
}

type CreateUserParams struct {
	Username  string
	Email     string
	Firstname string
	Lastname  string
	Password  string
}
