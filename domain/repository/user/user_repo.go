package repository

// UserProfileRepo - 登録ユーザー情報操作用のリポジトリ
type UserProfileRepo interface {
	Get()
	Insert()
	Update()
	Delete()
}

// UserImageRepo - 登録ユーザープロフィール画像操作用のリポジトリ
type UserImageRepo interface {
	Get()
	Insert()
	Update()
	Delete()
}
