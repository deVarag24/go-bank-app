package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

// store provides all func to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// create new store
func NewStore(db * sql.DB) *Store{
	return &Store{
		db: db,
		Queries: New(db),
	}
}

// to execute func within a db transaction
func (store *Store) executeTx(ctx context.Context, fn func(*Queries) error) error{
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil{
		return err
	}
	q := New(tx)

	err = fn(q)

	if err != nil{
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
	}
	return tx.Commit()
}

// contains input parameter for ransfer tx
type TransferTxParams struct{
	FromAccountId uuid.UUID
	ToAccountId uuid.UUID
	Amount int64
}

// result of transfer tx
type TransferTxResult struct{
	Transfer Transfer
	FromAccount Account
	ToAccount Account
	FromEntry Entry
	ToEntry Entry
}

// it performs money transfer from one account to other in one tx
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error){
	var result TransferTxResult

	err := store.executeTx(ctx, func(q *Queries) error{
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID: arg.ToAccountId,
			Amount: arg.Amount,
			ID: uuid.New(),
		})

		if err != nil{
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			ID: uuid.New(),
			AccountID: arg.FromAccountId,
			Amount: -arg.Amount,
		})

		if err != nil{
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			ID: uuid.New(),
			AccountID: arg.ToAccountId,
			Amount: arg.Amount,
		})

		if err != nil{
			return err
		}

		// TODO: update accounts

		return nil
	})
	 return result, err
}

