package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"lottery/bootstrap"
	"lottery/services"
	"lottery/web/controllers"
	"lottery/web/middleware"
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

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(
		userService,
		codeService,
		blackipService,
		giftService,
		resultService,
		userDayService,
	)
	admin.Handle(new(controllers.AdminController))

	adminGift := admin.Party("/gift")
	adminGift.Router.Use(middleware.BasicAuth)
	adminGift.Register(
		giftService,
	)
	adminGift.Handle(new(controllers.AdminGiftController))

	adminCode := admin.Party("/code")
	adminCode.Router.Use(middleware.BasicAuth)
	adminCode.Register(
		codeService,
	)
	adminGift.Handle(new(controllers.AdminCodeController))
}
