package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"lottery/bootstrap"
	"lottery/services"
	"lottery/web/controllers"
)

func Configure(b *bootstrap.Bootstrapper) {
	userService := services.NewUserService()
	codeService := services.NewCodeService()
	blackipService := services.NewBlackipService()
	giftService := services.NewGiftService()
	resultService := services.NewResultService()
	userDayService := services.NewUserDayService()

	index := mvc.New(b.Party("/"))
	index.Register(
		userService,
		codeService,
		blackipService,
		giftService,
		resultService,
		userDayService,
	)
	index.Handle(new(controllers.IndexController))
}
