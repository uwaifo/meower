package db

import (
	"context"

	"github.com/uwaifo/meower/schema"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertRow(ctx context.Context, meow schema.Meow) error {
	return impl.InsertMeow(ctx, meow)
}

func ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error) {
	return impl.ListMeows(ctx, skip, take)
}

/*
This is a straightforward way of achieving inversion of control.
By using Repository interface you allow any concrete implementation to be injected at runtime,
and all function calls will be delegated to the impl object.


*/
