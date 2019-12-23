package usecase

import (
	"fmt"

	"github.com/YuichiKadota/introther/domain/model"
	repository "github.com/YuichiKadota/introther/domain/repository/user"
)

type UeserUsecsse struct {
	userRepo repository.UserProfileRepo
}

func NewUserUsecase(userRepo repository.UserProfileRepo) UeserUsecsse {
	var u UeserUsecsse

	u = UeserUsecsse{userRepo: userRepo}
	return u
}

func (u *UeserUsecsse) Register(model.User) (model.User, error) {

	var user model.User
	reuser, err := u.Register(user)

	if err != nil {
		err = fmt.Errorf("ユーザー登録に失敗しました。 %w", err)
		return reuser, err
	}

	return reuser, nil
}
