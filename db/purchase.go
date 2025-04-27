package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/luke_design_pattern/util"
)

// {
// 	"PURCHASE_HISTORIES": [
// 		{
// 			"ID" : 3,
// 			"TRANSACTION_NUMBER" : "TRXPRC_20250421_RANDOMCHAR1"
// 			"DATE_OF_SALE" : "2024-12-17T16:20:18.000Z",
// 			"CUSTOMER_ID" : 4,
// 			"TOTAL_PRICE_PAYMENT" : 110000,
// 			"STATUS" : "completed",
// 		}
// 	]}

// {
// 	"PURCHASE_ITEMS": [
// 		{
// 			"ID" : 6,
// 			"BOOK_ID" : 93,
// 			"PURCHASE_HISTORY_ID" : 2,
// 			"QTY" : 2,
// 			"TOTAL_PRICE" : 0
// 		}
// 	]}

// create temporary purchase history
type BookModel struct {
	BookID   int
	Title    string
	Price    float64
	StockQty int
	AuthorID int
}

type PurchaseBook struct {
	Date              *time.Time
	BookID            int
	PurchaseHistoryID int
	Qty               int
	TotalPrice        int
}

type CreatePurchaseBookParams struct {
	// Date               *time.Time ==> auto generate from golang time.Date
	BookID            int
	PurchaseHistoryID int
	Qty               int
	TotalPrice        int
}

type PurchaseHistory struct {
	PurchaseID        int
	Date              *time.Time
	CustomerID        int
	TotalPricePayment float64
	Status            string
	TransactionNumber string
}

type CreatePurchaseHistoryParams struct {
	// Date              *time.Time ==> auto generate from golang time.Date
	CustomerID        int
	TotalPricePayment float64
	Status            string // pending or completed
	TransactionNumber string // PRCBOOK_20250421_RANDOMCHAR
}

func (cphp *CreatePurchaseHistoryParams) GenerateRandomTrxNumber() string {
	year, month, date := time.Now().Date()

	month_int := int(month)
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	scd := time.Now().Second() * int(util.RandomInt(0, 5000))
	custid := cphp.CustomerID
	generateTrxNumber := fmt.Sprintf("PRCBOOK%d%d%d%d%d%d%d", year, month_int, date, hour, minute, scd, custid)

	return generateTrxNumber
}

type CreatePurchaseHistoryResult struct {
}

type PurchaseItem struct {
	PurchaseItemID    int
	BookID            int
	PurchaseHistoryID int
	Qty               int
	TotalPrice        float64
}

type BookToPurchase struct {
	BookID            int
	Qty               int
	PurchaseHistoryID int
	CurrentStockQty   int
}

type ListBooksToPurchase []BookToPurchase

type CreateBookToPurchaseParams struct {
	BookID            int
	PurchaseHistoryID int
	Qty               int
	TotalPrice        float64
}

type CreatePurchaseBookTxParams struct{}
type CreatePurchaseBookTxResult struct{}

// CreatePurchase() PurchaseHistory => purchase history
// AddListBook()
// DeletePurchase() ["delete list to buy", "delete purchase history"]
// FinalizePurchase() => update purchase history "completed", set "total_price_payment"

const createPurchaseHistory = `
INSERT INTO PURCHASE_HISTORIES (
	DATE_OF_SALE, 
	CUSTOMER_ID, 
	TOTAL_PRICE_PAYMENT, 
	"STATUS", 
	TRANSACTION_NUMBER
) VALUES (CURRENT_TIMESTAMP, :1, :2, :3, :4)
	RETURNING ID, TRANSACTION_NUMBER INTO :5, :6
`

func (q *Queries) CreatePurchaseHistory(ctx context.Context, arg CreatePurchaseHistoryParams) (PurchaseHistory, error) {

	var i PurchaseHistory

	_, err := q.db.ExecContext(ctx, createPurchaseHistory,
		arg.CustomerID,
		arg.TotalPricePayment,
		arg.Status,
		arg.TransactionNumber,
		sql.Out{Dest: &i.PurchaseID},
		sql.Out{Dest: &i.TransactionNumber},
	)

	return i, err
}

const fetchBook = `
	SELECT 
		b.ID, 
		b.STOCK_QTY, 
		b.TITLE,
		b.PRICE 
	FROM BOOKS b WHERE b.ID = :1
`
const createNewPurchaseItemsBook = `
 INSERT INTO PURCHASE_ITEMS 
	(BOOK_ID, PURCHASE_HISTORY_ID, QTY, TOTAL_PRICE) 
 VALUES(:1, :2, :3, :4)
`

func (q *Queries) AddListBook(ctx context.Context, arg CreateBookToPurchaseParams) (BookToPurchase, error) {

	// note ==> harusnya dalam transaksi
	var i BookToPurchase
	var bookModel BookModel
	err := q.db.QueryRowContext(ctx, fetchBook, arg.BookID).Scan(
		&bookModel.BookID,
		&bookModel.StockQty,
		&bookModel.Title,
		&bookModel.Price,
	)

	i.BookID = bookModel.BookID
	i.Qty = arg.Qty
	i.PurchaseHistoryID = arg.PurchaseHistoryID

	if err != nil {
		if err == sql.ErrNoRows {
			return i, fmt.Errorf("ID buku tersebut tidak terdaftar atau tidak ada ==> %d", arg.BookID)
		}

		return i, err
	}

	i.CurrentStockQty = bookModel.StockQty

	if bookModel.StockQty < 1 {
		return i, fmt.Errorf("gagal, stok buku %s habis", bookModel.Title)
	}

	// check kuota < permintaan
	if bookModel.StockQty < arg.Qty {
		return i, fmt.Errorf("gagal, jumlah pembelian buku melebihi stok persediaan")
	}

	totalPricePerBook := bookModel.Price * float64(arg.Qty)

	_, err = q.db.ExecContext(ctx, createNewPurchaseItemsBook,
		arg.BookID,
		arg.PurchaseHistoryID,
		arg.Qty,
		totalPricePerBook,
	)
	if err != nil {
		// if strings.Contains(err.Error(), "ORA-02091") {
		// 	log.Printf("1: %v", err)
		// 	// return i, fmt.Errorf("ORA-02091: foreign key violation: purchase_history_id %d not found", arg.PurchaseHistoryID)
		// 	return i, err
		// }
		if strings.Contains(err.Error(), "ORA-02291") {
			log.Printf("2: %v", err)
			// return i, fmt.Errorf("ORA-02091: foreign key violation: purchase_history_id %d not found", arg.PurchaseHistoryID)
			return i, err
		}
		return i, err
	}

	return i, nil
}

func (q *Queries) DeletePurchase(ctx context.Context, arg ...interface{}) error {
	return nil
}

func (q *Queries) FinalizePurchase() error {
	return nil
}

func (store *SQLStore) PurchaseBookTx(ctx context.Context, arg CreatePurchaseBookTxParams) (CreatePurchaseBookTxResult, error) {
	return CreatePurchaseBookTxResult{}, nil
}
