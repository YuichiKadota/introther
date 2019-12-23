package model

import (
	"time"
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
