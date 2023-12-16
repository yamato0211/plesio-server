package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
	"github.com/yamato0211/plesio-server/pkg/utils/uuid"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

// idによるuserの単一取得
func (ur *userRepository) Select(ctx echo.Context, id string) (*entity.User, error) {
	sql := `SELECT * FROM users WHERE id = ?`
	user := entity.User{}
	err := ur.db.Get(&user, sql, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) Insert(ctx echo.Context, name string, email string, git_id string) error {
	sql := `INSERT INTO users (id, name, email, git_id) VALUES (:id, :name, :email, :git_id);`
	in := entity.User{
		ID:    uuid.NewUUID(),
		Name:  name,
		Email: email,
		GitID: git_id,
	}
	_, err := ur.db.NamedExec(sql, in)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) LoginBonus(ctx echo.Context, id string) (*entity.User, error) {
	sql := `SELECT * FROM users WHERE id = ?`
	user := entity.User{}
	err := ur.db.Get(&user, sql, id)

	if err != nil {
		return nil, err
	} else if user.IsLogined == false {
		sql := `UPDATE users SET is_logined = true WHERE id = ?`
		_, err := ur.db.Exec(sql, id)
		if err != nil {
			return nil, err
		} else {
			sql := `UPDATE users SET coin = coin + 1 WHERE id = ?`
			_, err = ur.db.Exec(sql, id)
			if err != nil {
				return nil, err
			}
			return &user, nil
		}
	} else {
		return nil, nil
	}
}
