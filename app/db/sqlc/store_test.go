package db

import (
	"context"
	"my-go-app/app/util"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T){

	store := NewStore(testDB)

	acc_1 := createRandomAccount(t)
	acc_2 := createRandomAccount(t)

	// run n concurrent transfer transactions
	n:=2

	amount:=util.RandomInt(100, 5000)

	errs:= make(chan error, n)
	results:= make(chan TransferTxResult, n)

	for i := 0; i < n; i++ {
		go func(fromID, toID uuid.UUID, amt int64) {
			ctx := context.Background()
			res, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountId: fromID,
				ToAccountId:   toID,
				Amount:        amt,
			})
			results <- res
			errs <- err
		}(acc_1.ID, acc_2.ID, amount)
	}
	

	for i:=0; i<n; i++{
		err:= <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, acc_1.ID, transfer.FromAccountID)
		require.Equal(t, acc_2.ID, transfer.ToAccountID)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)


		// check entry

		from_entry := result.FromEntry

		require.NotEmpty(t, from_entry)
		require.Equal(t, acc_1.ID, from_entry.AccountID)
		require.Equal(t, -amount, from_entry.Amount)

		_, err = store.GetEntry(context.Background(), from_entry.ID)
		require.NoError(t, err)


		to_entry := result.ToEntry

		require.NotEmpty(t, to_entry)
		require.Equal(t, acc_2.ID, to_entry.AccountID)
		require.Equal(t, amount, to_entry.Amount)

		_, err = store.GetEntry(context.Background(), to_entry.ID)
		require.NoError(t, err)

		// check account
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, fromAccount.ID, acc_1.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, toAccount.ID, acc_2.ID)


		// check account balance
		dif_1:= acc_1.Balance - fromAccount.Balance
		dif_2:= toAccount.Balance - acc_2.Balance

		require.Equal(t, dif_1, dif_2)
		require.True(t, dif_1 > 0)
	}

	updated_acc_1, err := testQueries.GetAccount(context.Background(), acc_1.ID)
	require.NoError(t, err)
	require.Equal(t, acc_1.Balance-int64(n)*amount, updated_acc_1.Balance)

	updated_acc_2, err := testQueries.GetAccount(context.Background(), acc_2.ID)
	require.NoError(t, err)
	require.Equal(t, acc_2.Balance+int64(n)*amount, updated_acc_2.Balance)


}