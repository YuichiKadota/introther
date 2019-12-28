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

func (u *UeserUsecsse) Register(user *model.User) (model.User, error) {

	// TODO パスワード暗号化処理を追加する
	user.SetInsertDate()
	user.SetUpdatedateDate()
	err := user.Validate()
	if err != nil {
		err = fmt.Errorf("入力項目が不適切です。 %w", err)
		return *user, err
	}

	gotUser, err := u.userRepo.Get(user.UserID)
	if err != nil {
		err = fmt.Errorf("ユーザーデータの取得に失敗しました。 %w", err)
		return *user, err
	}

	err = user.DuplicateUserCheck(*gotUser)
	if err != nil {
		return *user, err
	}

	reuser, err := u.userRepo.Insert(user)

	if err != nil {
		err = fmt.Errorf("ユーザー登録に失敗しました。 %w", err)
		return *reuser, err
	}

	return *reuser, nil
}
