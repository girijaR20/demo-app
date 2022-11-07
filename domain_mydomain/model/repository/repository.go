package repository

import (
	"context"
	"demo-app/domain_mydomain/model/entity"
)

type SaveUserRepo interface {
	SaveUser(ctx context.Context, obj *entity.User) error
	//InvalidUser(ctx context.Context, obj*entity.User)error
}
