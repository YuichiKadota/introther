package model

import (
	"time"
)

// User - 登録ユーザーモデル
type User struct {
	UserID     string `validate:"required"`
	Password   string `validate:"required"`
	NickName   string `validate:"required"`
	Profile    string `validate:"gte=0,lt=130"` // 0以上、130未満
	ImageURL   string `validate:"url"`
	InsertDate time.Time
	UpdateDate time.Time
}
