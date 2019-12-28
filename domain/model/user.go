package model

import (
	"fmt"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/go-ozzo/ozzo-validation/v3/is"
)

// User - 登録ユーザーモデル
type User struct {
	UserID     string
	Password   string
	NickName   string
	Profile    string
	ImageURL   string
	InsertDate time.Time
	UpdateDate time.Time
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		// UserIDは空を許容せず、5から20までの長さ、半角英数字記号のみ
		validation.Field(&u.UserID,
			validation.Required,
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9!-/:-@¥[-`{-~]+$")),
			validation.RuneLength(5, 20)),
		// Passwordは空を許容せず、5から20までの長さ、半角英数字記号のみ
		validation.Field(&u.Password,
			validation.Required,
			validation.RuneLength(5, 20),
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9!-/:-@¥[-`{-~]+$"))),
		// NickNameは空を許容せず、5から20までの長さ、半角英数字記号のみ
		validation.Field(&u.NickName,
			validation.Required,
			validation.RuneLength(5, 20),
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9!-/:-@¥[-`{-~]+$"))),
		// Profileは0~255の文字列
		validation.Field(&u.Profile,
			validation.RuneLength(0, 255)),
		// ImageURLはURL
		validation.Field(&u.ImageURL,
			is.URL),
	)
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
