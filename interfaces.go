package icrud

import "context"

type Creatable[T any, U any] interface {
	Insert(ctx context.Context, n T) (*U, error)
}

type Readable[T any] interface {
	List(ctx context.Context) ([]T, error)
	Get(ctx context.Context, id uint64) (*T, error)
}

type Updatable[T any] interface {
	Update(ctx context.Context, n T) (*T, error)
}

type Deletable interface {
	Delete(ctx context.Context, id uint64) error
}
