package routes

import (
	"bifromq_engine/pkg/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	//r := gin.Default()
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("captch"))))
	// Register middlewares
	r.Use(Logger(), Authentication(), Cors())
	// Register routes
	r.GET("/ping", api.Ping)

	r.POST("/api/auth/login", api.Auth.Login)
	r.GET("/api/auth/captcha", api.Auth.Captcha)

	r.POST("/api/auth/logout", api.Auth.Logout)
	r.POST("/api/auth/password", api.Auth.Logout)
	//
	r.GET("/api/user", api.User.List)
	r.POST("/api/user", api.User.Add)
	r.DELETE("/api/user/:id", api.User.Delete)
	r.PATCH("/api/user/password/reset/:id", api.User.Update)
	r.PATCH("/api/user/:id", api.User.Update)
	r.PATCH("/api/user/profile/:id", api.User.Profile)
	r.GET("/api/user/detail", api.User.Detail)

	r.GET("/api/role", api.Role.List)
	r.POST("/api/role", api.Role.Add)
	r.PATCH("/api/role/:id", api.Role.Update)
	r.DELETE("/api/role/:id", api.Role.Delete)
	r.PATCH("/api/role/users/add/:id", api.Role.AddUser)
	r.PATCH("/api/role/users/remove/:id", api.Role.RemoveUser)
	r.GET("/api/role/page", api.Role.ListPage)
	r.GET("/api/role/permissions/tree", api.Role.PermissionsTree)

	r.POST("/api/permission", api.Permissions.Add)
	r.PATCH("/api/permission/:id", api.Permissions.PatchPermission)
	r.DELETE("/api/permission/:id", api.Permissions.Delete)
	r.GET("/api/permission/tree", api.Permissions.List)
	r.GET("/api/permission/menu/tree", api.Permissions.List)
	r.GET("/api/permission/menu/validate", api.Permissions.ValidateMenu)

	r.GET("/api/node", api.Worker.List)
	r.POST("/api/node/register", api.Worker.Register)
	r.GET("/api/configuration", api.Worker.GetConfiguration)
	r.PUT("/api/configuration", api.Worker.UpdateConfiguration)
	r.GET("/api/ruleset", api.RuleSet.List)
	r.POST("/api/ruleset", api.RuleSet.Add)
	r.GET("/api/ruleset/:id", api.RuleSet.Detail)
}
