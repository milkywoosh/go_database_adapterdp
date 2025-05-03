package db

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreatePurchaseHistory(t *testing.T) PurchaseHistory {

	arg := CreatePurchaseHistoryParams{
		CustomerID:        3,
		TotalPricePayment: 10000.5,
		Status:            "pending",
		PurchaseNumber:    GenerateRandomTrxNumber(4),
	}

	result, err := testStoreOra.CreatePurchaseHistory(context.Background(), arg)
	log.Printf("arg => %v", arg)
	log.Printf("result => %v", result)
	// log.Printf("check log error ==> %v", err)

	require.NoError(t, err, "check error create purchase history")
	require.NotEmpty(t, result)
	require.Positive(t, result.PurchaseID, "generated Purchase ID")
	require.Equal(t, arg.PurchaseNumber, result.PurchaseNumber, "check generated transaction number")

	return result
}

func TestCreatePurchaseHistory(t *testing.T) {
	trx_number := GenerateRandomTrxNumber(4)
	arg := CreatePurchaseHistoryParams{
		CustomerID:        3,
		TotalPricePayment: 0.0,
		Status:            "pending",
		PurchaseNumber:    trx_number,
	}

	result, err := testStorePG.CreatePurchaseHistory(context.Background(), arg)
	log.Printf("arg => %v", arg)
	log.Printf("result => %v", result)
	log.Printf("check log error ==> %v", err)

	require.NoError(t, err, "check error create purchase history")
	require.NotEmpty(t, result)
	require.Positive(t, result.PurchaseID, "generated Purchase ID")
	require.Equal(t, result.PurchaseNumber, trx_number, "check generated transaction number")
}

func TestAddListBook(t *testing.T) {
	// purchaseHistoryID := CreatePurchaseHistory(t).PurchaseID
	purchaseHistoryID := 4

	arg := CreateBookToPurchaseParams{
		BookID:            3000,
		PurchaseHistoryID: purchaseHistoryID,
		Qty:               101,
		TotalPrice:        0.0,
		PurchaseNumber:    "PRCBOOK2025531955112124",
	}

	// var bookToPurchase BookToPurchase
	var err error
	_, err = testStorePG.AddListBook(context.Background(), arg)
	// log.Printf("curr qty ==> %v", bookToPurchase.CurrentStockQty)
	// log.Printf("arg qty ==> %v", bookToPurchase.Qty)

	log.Printf("err %v", err)
	// require.NoError(t, err)
	require.False(t, err == ErrIDBukuTidakTerdaftar, fmt.Sprintf("%s ==> %s", ErrIDBukuTidakTerdaftar, "buku tidak terdaftar, buku harus terdaftar di database"))
	require.False(t, err == ErrStokBukuHabis, fmt.Sprintf("%s ==> %s", ErrStokBukuHabis, "stok buku harus lebih dari 0"))
	require.False(t, err == ErrStokBukuKurang, fmt.Sprintf("%s ==> %s", ErrStokBukuKurang, "stok buku saat ini lebih sedikit dari penawaran"))
	// require.True(t, bookToPurchase.CurrentStockQty > 1, "stock harus lebih besar dari 0")
	// require.True(t, bookToPurchase.CurrentStockQty >= arg.Qty, "jml stock saat ini harus lebih besar dari jumlah permintaan")

}

func TestEditListBook(t *testing.T) {

	price := 1000.0
	args := EditBookToPurchaseParams{
		Qty:               6,
		TotalPrice:        6 * price,
		BookID:            2800,
		PurchaseHistoryID: 4,
		PurchaseNumber:    "PRCBOOK2025531955112124",
	}

	// log.Printf("bookid %d, prchistiD %d, prcNumber %s", args.BookID, args.PurchaseHistoryID, args.PurchaseNumber)

	updatedRow, err := testStorePG.EditListBookTx(context.Background(), args)

	require.True(t, updatedRow == 1, fmt.Sprintf("data terupdate => %d . update data harus 1 row, tidak boleh lebih atau kurang", updatedRow))
	require.False(t, updatedRow == 0, "tidak boleh update 0 row")
	require.False(t, ErrUpdateNolData != nil, ErrUpdateNolData)
	require.NoError(t, err, "check error apapun terakhir")
}

func TestAdjustSTockBook(t *testing.T) {
	bookID := 6
	corrector := -20

	err := testStorePG.AdjustStockBook(context.Background(), bookID, corrector)
	log.Printf("%v", ErrIDBukuTidakTerdaftar)
	require.NoError(t, err)
	require.True(t, ErrIDBukuTidakTerdaftar == nil, "book is existed is NO")
}
