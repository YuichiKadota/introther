package usecase

import (
	"fmt"

	"github.com/YuichiKadota/introther/domain/model"
	repository "github.com/YuichiKadota/introther/domain/repository/post"
)

//ユースケースを構造体として定義
type PostUseCase struct {
	postRepository repository.PostRepo
}

//Postに対するUsecaseを作成
func NewPostUseCase(pr repository.PostRepo) PostUseCase {
	var pu PostUseCase
	pu = PostUseCase{postRepository: pr}
	return pu
}

//投稿登録メソッドを作成
func (pu *PostUseCase) Submit(post *model.Post) (model.Post, error) {

	//投稿ID、登録日時、更新日時を生成
	post.SetPostId()
	post.SetInsertDate()
	post.SetUpdateDate()

	//入力値チェック
	err := post.Validate()
	if err != nil {
		err := fmt.Errorf("入力内容が不適切です。:%w", err)
		return *post, err
	}

	//投稿重複チェック
	err = post.DuplicatePostCheck(*post)
	if err != nil {
		return *post, err
	}

	//投稿の登録
	repo, err := pu.postRepository.Insert(post)

	if err != nil {
		err := fmt.Errorf("投稿に失敗しました: %w", err)
		return *repo, err
	}

	return *repo, nil
}
