package handler

import (
	"net/http"

	"github.com/YuichiKadota/introther/domain/model"
	"github.com/YuichiKadota/introther/usecase"
	"github.com/labstack/echo"
)

type UserHandler struct {
	userUsecase usecase.UeserUsecsse
}

func NewUsersHandler(userUsecase usecase.UeserUsecsse) UserHandler {
	userHandler := UserHandler{userUsecase: userUsecase}
	return userHandler
}

func (h *UserHandler) View() echo.HandlerFunc {
	return nil
}

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

func (h *UserHandler) Edit() echo.HandlerFunc {
	return nil
}
func (h *UserHandler) Delete() echo.HandlerFunc {
	return nil
}
