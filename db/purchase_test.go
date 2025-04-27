package db

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreatePurchaseHistory(t *testing.T) PurchaseHistory {

	arg := CreatePurchaseHistoryParams{
		CustomerID:        3,
		TotalPricePayment: 10000.5,
		Status:            "pending",
		TransactionNumber: "",
	}
	generateTrxNumber := arg.GenerateRandomTrxNumber()
	arg.TransactionNumber = generateTrxNumber

	result, err := testStoreOra.CreatePurchaseHistory(context.Background(), arg)
	log.Printf("arg => %v", arg)
	log.Printf("result => %v", result)
	// log.Printf("check log error ==> %v", err)

	require.NoError(t, err, "check error create purchase history")
	require.NotEmpty(t, result)
	require.Positive(t, result.PurchaseID, "generated Purchase ID")
	require.Equal(t, arg.TransactionNumber, result.TransactionNumber, "check generated transaction number")

	return result
}

func TestAddListBook(t *testing.T) {
	// purchaseHistoryID := CreatePurchaseHistory(t).PurchaseID
	purchaseHistoryID := 1090

	arg := CreateBookToPurchaseParams{
		BookID:            471,
		PurchaseHistoryID: purchaseHistoryID,
		Qty:               100,
		TotalPrice:        0.0,
	}

	var bookToPurchase BookToPurchase
	var err error
	bookToPurchase, err = testStoreOra.AddListBook(context.Background(), arg)
	// log.Printf("curr qty ==> %v", bookToPurchase.CurrentStockQty)
	// log.Printf("arg qty ==> %v", bookToPurchase.Qty)

	require.Contains(t, err.Error(), "ORA-02091", "error violation integrity constraint")
	require.NoError(t, err)
	require.True(t, bookToPurchase.CurrentStockQty > 1, "stock harus lebih besar dari 0")
	require.True(t, bookToPurchase.CurrentStockQty >= arg.Qty, "jml stock saat ini harus lebih besar dari jumlah permintaan")

}
