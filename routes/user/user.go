package routes

import (
	"user-service/controllers"
	"user-service/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
}

type IUserRoute interface {
	Run()
}

func NewUserRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup) IUserRoute {
	return &UserRoute{
		controller: controller,
		group:      group,
	}
}

func (u *UserRoute) Run() {
	group := u.group.Group("/auth")
	group.GET("/user", middlewares.Authenticated(), u.controller.GetUserController().GetUserLogin)
	group.GET("/:uuid", middlewares.Authenticated(), u.controller.GetUserController().GetUserByUUID)
	group.PUT("/:uuid", middlewares.Authenticated(), u.controller.GetUserController().Update)

	group.POST("/register", u.controller.GetUserController().Register)
	group.POST("/login", u.controller.GetUserController().Login)
}
