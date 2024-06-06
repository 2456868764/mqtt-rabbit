package resp

import (
	"time"

	"bifromq_engine/pkg/model/entity"
)

type LoginRes struct {
	AccessToken string `json:"accessToken"`
}

type UserDetailRes struct {
	entity.User
	Profile     *entity.Profile `json:"profile"`
	Roles       []*entity.Role  `json:"roles" `
	CurrentRole *entity.Role    `json:"currentRole"`
}

type RoleListRes []*entity.Role

type UserListItem struct {
	ID         int            `json:"id"`
	Username   string         `json:"username"`
	Enable     bool           `json:"enable"`
	CreateTime time.Time      `json:"createTime"`
	UpdateTime time.Time      `json:"updateTime"`
	Gender     int            `json:"gender"`
	Avatar     string         `json:"avatar"`
	Address    string         `json:"address"`
	Email      string         `json:"email"`
	Roles      []*entity.Role `json:"roles"`
}

type UserListRes struct {
	PageData []UserListItem `json:"pageData"`
	Total    int64          `json:"total"`
}
type RoleListPageItem struct {
	entity.Role
	PermissionIds []int64 `json:"permissionIds" gorm:"-"`
}
type RoleListPageRes struct {
	PageData []RoleListPageItem `json:"pageData"`
	Total    int64              `json:"total"`
}
