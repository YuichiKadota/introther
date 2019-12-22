package presenter

import (
	"github.com/YuichiKadota/introther/presenter/handler"
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, userHandler handler.UserHandler) {

	e.GET("/userProfile", userHandler.View())

	e.POST("/userRegister", userHandler.Register())

	e.POST("/userEdit", userHandler.Edit())

	e.POST("/userDelete", userHandler.Delete())

}
