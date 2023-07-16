package db

import (
	"context"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	//run n concurrent operations
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TranferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.Transferx(context.Background(), TransfersTxParams{
				FromAccountId: account1.ID,
				ToAccountId:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()

	}
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
		//check transfer
		transfer := result.Transer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, errs := store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, errs)
		//check from Entry
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		// get entry
		__, errs := store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, errs)
		require.NotEmpty(t, __)

		//check to Entry
		ToEntry := result.ToEntry
		require.NotEmpty(t, ToEntry)
		require.Equal(t, account2.ID, ToEntry.AccountID)
		require.Equal(t, amount, ToEntry.Amount)
		require.NotZero(t, ToEntry.ID)
		require.NotZero(t, ToEntry.CreatedAt)
		// get entry
		___, errs := store.GetEntry(context.Background(), ToEntry.ID)
		require.NoError(t, errs)
		require.NotEmpty(t, ___)

	}
}
