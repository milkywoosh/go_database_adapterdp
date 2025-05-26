package purchase

import "time"

type CreatePurchaseBookTxParams struct{}
type CreatePurchaseBookTxResult struct{}
type PurchaseBook struct {
	Date              *time.Time
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
	PurchaseNumber    string
}

type PurchaseItem struct {
	PurchaseItemID,
	BookID,
	PurchaseHistoryID,
	Qty int
	TotalPrice float64
}

type Book struct {
	BookID   int
	Title    string
	Price    float64
	StockQty int
	AuthorID int
}
