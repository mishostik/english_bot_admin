package user

import "context"

type Repository interface {
	Select(ctx context.Context) ([]User, error)
}
