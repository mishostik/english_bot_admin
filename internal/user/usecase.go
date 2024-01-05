package user

import "context"

type Usecase interface {
	GetAll(context_ context.Context) ([]User, error)
}
