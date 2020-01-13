package repository

import (
	"github.com/YuichiKadota/introther/domain/model"
)

// PostRepo - 投稿操作用のリポジトリ
type PostRepo interface {

	// TODO 更新、削除
	Get(string) (*model.Post, error)
	Insert(*model.Post) (*model.Post, error)
	//	Update(*model.Post) (*model.Post, error)
	//	Delete(*model.Post) (bool, error)
}
