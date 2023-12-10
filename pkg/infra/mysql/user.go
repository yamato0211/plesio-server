package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

// userの単一取得
func (ur *userRepository) Select(ctx echo.Context, id string) (*entity.User, error) {
	var user *entity.User
	err := ur.db.Get(user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
