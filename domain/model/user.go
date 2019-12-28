package model

import (
	"fmt"
	"time"
)

// User - 登録ユーザーモデル
type User struct {
	UserID     string    `validate:"required"`
	Password   string    `validate:"required"`
	NickName   string    `validate:"required"`
	Profile    string    `validate:"gte=0,lt=130"` // 0以上、130未満
	ImageURL   string    `validate:"url"`
	InsertDate time.Time `validate:"lte"`
	UpdateDate time.Time `validate:"lte"`
}

// DuplicateUserCheck - ユーザー登録時は既存のユーザーIDとの重複は許さない
func (u *User) DuplicateUserCheck(existUser User) error {

	if u.UserID == existUser.UserID {
		return fmt.Errorf("そのユーザーIDは既に登録されています。")
	}
	return nil
}

func (u *User) SetInsertDate() {
	u.InsertDate = time.Now()
}

func (u *User) SetUpdatedateDate() {
	u.UpdateDate = time.Now()
}
