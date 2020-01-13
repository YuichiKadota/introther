package model

import (
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/go-ozzo/ozzo-validation/v3/is"
	"github.com/google/uuid"
)

// Post - 投稿モデル
type Post struct {
	PostID       string
	PostUserID   string
	PostedUserID string
	Text         string
	//TODO 複数画像を伴う投稿
	ImageURL   string
	InsertDate time.Time
	UpdateDate time.Time
}

func (p *Post) Validate() error {
	return validation.ValidateStruct(p,
		// PostTextはo～255の文字列
		validation.Field(&p.Text,
			validation.RuneLength(0, 255)),
		// ImageURLはURL
		validation.Field(&p.ImageURL,
			is.URL),
	)
}

// DuplicateUserCheck - ユーザXからユーザYへの複数の投稿を許さない
func (p *Post) DuplicatePostCheck(existPost Post) error {

	if p.PostUserID == existPost.PostUserID && p.PostedUserID == existPost.PostedUserID {
		return fmt.Errorf("すでに紹介済みです。")
	}
	return nil
}

func (p *Post) SetPostId() {
	//UUIDの発行
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	//UUIDをpostIDに代入
	p.PostID = u.String()
}

func (p *Post) SetInsertDate() {
	p.InsertDate = time.Now()
}

func (p *Post) SetUpdateDate() {
	p.UpdateDate = time.Now()
}
