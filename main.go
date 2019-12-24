package main

import (
	"fmt"

	"github.com/YuichiKadota/introther/di"
	"github.com/YuichiKadota/introther/presenter"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	fmt.Println("sever start")
	userHandler, err := di.InjectUserHandler()
	if err != nil {
		log.Fatalf("ユーザー処理関連の依存解決に失敗しました。 %v", err)
	}
	e := echo.New()
	e.Validator = presenter.NewValidator()
	presenter.InitRouting(e, userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
