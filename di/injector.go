package di

import (
	"fmt"

	repository "github.com/YuichiKadota/introther/domain/repository/user"
	infra "github.com/YuichiKadota/introther/infra/user"
	"github.com/YuichiKadota/introther/presenter/handler"
	"github.com/YuichiKadota/introther/usecase"
)

func InjectUserRepository() (repository.UserProfileRepo, error) {
	dynamodbRepoImpl, err := infra.NewDynamoDBRepoImpl()

	if err != nil {
		err = fmt.Errorf("DynamoDBRepoImplの初期化に失敗しました。 %w", err)
		return nil, err
	}
	return dynamodbRepoImpl, nil
}

func InjectUserUsecase() (usecase.UeserUsecsse, error) {
	var userUsecase usecase.UeserUsecsse
	UserRepo, err := InjectUserRepository()

	if err != nil {
		err = fmt.Errorf("DynamoDBRepoImplの初期化に失敗しました。 %w", err)
		return userUsecase, err
	}

	userUsecase = usecase.NewUserUsecase(UserRepo)

	return userUsecase, nil
}

func InjectUserHandler() (handler.UserHandler, error) {

	var userHandler handler.UserHandler

	userUsecase, err := InjectUserUsecase()
	if err != nil {
		err = fmt.Errorf("userUsecaseの初期化に失敗しました。 %w", err)
		return userHandler, err
	}

	userHandler = handler.NewUsersHandler(userUsecase)

	return userHandler, nil
}
