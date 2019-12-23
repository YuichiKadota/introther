package handler

import (
	"net/http"

	"github.com/YuichiKadota/introther/domain/model"
	"github.com/YuichiKadota/introther/usecase"
	"github.com/labstack/echo"
)

// UserHandler - ユーザー登録のユースケースを呼び出す実装メソッドをもつ構造体
type UserHandler struct {
	userUsecase usecase.UeserUsecsse
}

// NewUsersHandler - ユーザー登録のユースケースを呼び出すための実装を返す
func NewUsersHandler(userUsecase usecase.UeserUsecsse) UserHandler {
	userHandler := UserHandler{userUsecase: userUsecase}
	return userHandler
}

// View - 単一のユーザー情報を引いてくるユースケースへハンドリングする（仮定義）
func (h *UserHandler) View() echo.HandlerFunc {
	return nil
}

// Register - ユーザー登録を行うユースケースへハンドリングする
func (h *UserHandler) Register() echo.HandlerFunc {

	return func(c echo.Context) error {
		var user model.User
		c.Bind(&user)
		reuser, err := h.userUsecase.Register(user)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, reuser)
	}

}

// Edit - ユーザー情報編集を行うユースケースへハンドリングする
func (h *UserHandler) Edit() echo.HandlerFunc {
	return nil
}

// Delete - ユーザー退会処理を行うユースケースへハンドリングする
func (h *UserHandler) Delete() echo.HandlerFunc {
	return nil
}
