package db

import "context"

type ctxKey struct{}

var txKey = ctxKey{}

func ContextWithTx(ctx context.Context, tx Store) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

func GetTx(ctx context.Context) (Store, bool) {
	tx, ok := ctx.Value(txKey).(Store)
	return tx, ok
}
