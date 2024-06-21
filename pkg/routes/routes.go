package routes

import (
	"bifromq_engine/pkg/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("captch"))))
	// Register middlewares
	r.Use(Logger(), Authentication(), Cors())
	// Register routes
	r.GET("/ping", api.Ping)

	r.POST("/auth/login", api.Auth.Login)
	r.GET("/auth/captcha", api.Auth.Captcha)

	r.POST("/auth/logout", api.Auth.Logout)
	r.POST("/auth/password", api.Auth.Logout)
	//
	r.GET("/user", api.User.List)
	r.POST("/user", api.User.Add)
	r.DELETE("/user/:id", api.User.Delete)
	r.PATCH("/user/password/reset/:id", api.User.Update)
	r.PATCH("/user/:id", api.User.Update)
	r.PATCH("/user/profile/:id", api.User.Profile)
	r.GET("/user/detail", api.User.Detail)

	r.GET("/role", api.Role.List)
	r.POST("/role", api.Role.Add)
	r.PATCH("/role/:id", api.Role.Update)
	r.DELETE("/role/:id", api.Role.Delete)
	r.PATCH("/role/users/add/:id", api.Role.AddUser)
	r.PATCH("/role/users/remove/:id", api.Role.RemoveUser)
	r.GET("/role/page", api.Role.ListPage)
	r.GET("/role/permissions/tree", api.Role.PermissionsTree)

	r.POST("/permission", api.Permissions.Add)
	r.PATCH("/permission/:id", api.Permissions.PatchPermission)
	r.DELETE("/permission/:id", api.Permissions.Delete)
	r.GET("/permission/tree", api.Permissions.List)
	r.GET("/permission/menu/tree", api.Permissions.List)

	r.GET("/node", api.Worker.List)
	r.POST("/node/register", api.Worker.Register)
	r.GET("/configuration", api.Worker.GetConfiguration)
	r.PUT("/configuration", api.Worker.UpdateConfiguration)
	return r
}
