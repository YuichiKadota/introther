package repository

import (
	"github.com/YuichiKadota/introther/domain/model"
)

// UserProfileRepo - 登録ユーザー情報操作用のリポジトリ
type UserProfileRepo interface {
	Get(string) (model.User, error)
	Insert(model.User) (model.User, error)
	Update(model.User) (model.User, error)
	Delete(model.User) (bool, error)
}
