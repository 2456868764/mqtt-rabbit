package entity

import (
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Enable     bool      `json:"enable"`
	CreateTime time.Time `json:"createTime" gorm:"column:createTime"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updateTime"`
}

func (User) TableName() string {
	return "user"
}

type UserRolesRole struct {
	UserId int `gorm:"column:userId"`
	RoleId int `gorm:"column:roleId"`
}

func (UserRolesRole) TableName() string {
	return "user_roles_role"
}

type Role struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Enable bool   `json:"enable"`
}

func (Role) TableName() string {
	return "role"
}

type RolePermissionsPermission struct {
	RoleId       int `gorm:"column:roleId"`
	PermissionId int `gorm:"column:permissionId"`
}

func (RolePermissionsPermission) TableName() string {
	return "role_permissions_permission"
}

type Profile struct {
	ID       int    `json:"id"`
	Gender   int    `json:"gender"`
	Avatar   string `json:"avatar"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	UserId   int    `gorm:"column:userId"`
	NickName string `gorm:"column:nickName"`
}

func (Profile) TableName() string {
	return "profile"
}

type Permission struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Code        string       `json:"code"`
	Type        string       `json:"type"`
	ParentId    *int         `json:"parentId" gorm:"column:parentId"`
	Path        string       `json:"path"`
	Redirect    string       `json:"redirect"`
	Icon        string       `json:"icon"`
	Component   string       `json:"component"`
	Layout      string       `json:"layout"`
	KeepAlive   int          `json:"keepAlive" gorm:"column:keepAlive"`
	Method      string       `json:"method"`
	Description string       `json:"description"`
	Show        int          `json:"show"`
	Enable      int          `json:"enable"`
	Order       int          `json:"order"`
	Children    []Permission `json:"children" gorm:"-"`
}

func (Permission) TableName() string {
	return "permission"
}
