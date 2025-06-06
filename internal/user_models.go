package internal

type Book struct {
	BookID   int
	Title    string
	Price    float64
	StockQty int
	AuthorID int
}

type Users struct {
	Username             string `json:"username"`
	Email                string `json:"email"`
	Firstname            string `json:"firstname"`
	Lastname             string `json:"lastname"`
	PasswordHistoryCount string `json:"password_history_count"`
}

type LoginUserTxParams struct {
	Username, RawPassword, EncryptedPW string
}

type LoginUserTxResult struct {
	Username, Password string
	IsSuccess          bool
	Msg                string
}

type UserCredential struct {
	Username, Password string
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

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user Users) error
}
