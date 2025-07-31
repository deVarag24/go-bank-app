package db

import (
	"context"
	"my-go-app/app/util"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account{
	arg := CreateAccountParams{
		Owner: util.RandomString(5),
		Balance: util.RandomInt(100, 10000),
		Currency: util.RandomEnum([]string{"USD", "INR"}),
		ID: util.RandomUuid(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.ID, account.ID)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	return account;
}

func deleteAccountById(t *testing.T, id uuid.UUID){
	err := testQueries.DeleteAccount(context.Background(), id)

	require.NoError(t, err)
}

func TestCreateAccount(t *testing.T){
	acc := createRandomAccount(t)

	defer deleteAccountById(t, acc.ID)
}

func TestDeleteAccount(t *testing.T){
	acc := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), acc.ID)

	require.NoError(t, err)
}


func TestGetAccount(t *testing.T){
	acc_1 := createRandomAccount(t)

	acc_2, err:= testQueries.GetAccount(context.Background(), acc_1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc_2)
	require.Equal(t, acc_1.ID, acc_2.ID)

	defer deleteAccountById(t, acc_1.ID)
}


func TestGetAllAccounts(t *testing.T){

	const no_of_accounts = 2
	var created_accounts []Account

	for i:=0; i<no_of_accounts; i++ {
		acc := createRandomAccount(t)
		created_accounts = append(created_accounts, acc)
	}

	arr:= GetAllAccountsParams{
		Limit: no_of_accounts,
		Offset: 0,
	}

	acc, err := testQueries.GetAllAccounts(context.Background(), arr)

	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, no_of_accounts, len(acc))

	for i:=0; i<no_of_accounts; i++{
		require.Equal(t, created_accounts[i].ID, acc[i].ID, "Account is not fetched")
	}

	defer func() {
		for _, acc := range created_accounts {
			deleteAccountById(t, acc.ID)
		}
	}()
}


func TestUpdateAccountBalance(t *testing.T){
	acc_1:= createRandomAccount(t)

	arg := UpdateAccountBalanceParams{
		ID: acc_1.ID,
		Balance: util.RandomInt(100, 5000),
	}

	acc_2, err := testQueries.UpdateAccountBalance(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, acc_2)
	require.Equal(t, arg.ID, acc_2.ID)
	require.Equal(t, arg.Balance, acc_2.Balance)

	defer deleteAccountById(t, acc_1.ID)

}

