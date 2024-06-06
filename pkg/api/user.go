package api

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

	"bifromq_engine/pkg/db"
	"bifromq_engine/pkg/model/entity"
	"bifromq_engine/pkg/model/req"
	"bifromq_engine/pkg/model/resp"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var User = &user{}

type user struct {
}

func (user) Detail(c *gin.Context) {
	var data = &resp.UserDetailRes{}
	var uid, _ = c.Get("uid")
	db.DB.Model(entity.User{}).Where("id=?", uid).Find(&data)
	db.DB.Model(entity.Profile{}).Where("userId=?", uid).Find(&data.Profile)
	urolIdList := db.DB.Model(entity.UserRolesRole{}).Where("userId=?", uid).Select("roleId")
	db.DB.Model(entity.Role{}).Where("id IN (?)", urolIdList).Find(&data.Roles)
	if len(data.Roles) > 0 {
		data.CurrentRole = data.Roles[0]
	}
	Success(c, data)
}

func (user) List(c *gin.Context) {
	var data = resp.UserListRes{
		PageData: make([]resp.UserListItem, 0),
	}
	var gender = c.DefaultQuery("gender", "")
	var enable = c.DefaultQuery("enable", "")
	var username = c.DefaultQuery("username", "")
	var pageNoReq = c.DefaultQuery("pageNo", "1")
	var pageSizeReq = c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(pageNoReq)
	pageSize, _ := strconv.Atoi(pageSizeReq)
	var profileList []entity.Profile
	orm := db.DB.Model(entity.Profile{})
	if gender != "" {
		orm = orm.Where("gender=?", gender)
	}
	if enable != "" {
		orm = orm.Where("userId in(?)", db.DB.Model(entity.User{}).Where("enable=?", enable).Select("id"))
	}
	if username != "" {
		orm = orm.Where("nickName like ?", "%"+username+"%")
	}

	orm.Count(&data.Total)
	orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&profileList)
	for _, datum := range profileList {
		var uinfo entity.User
		db.DB.Model(entity.User{}).Where("id=?", datum.UserId).First(&uinfo)
		var rols []*entity.Role
		db.DB.Model(entity.Role{}).Where("id IN (?)", db.DB.Model(entity.UserRolesRole{}).Where("userId=?", datum.UserId).Select("roleId")).Find(&rols)
		data.PageData = append(data.PageData, resp.UserListItem{
			ID:         uinfo.ID,
			Username:   uinfo.Username,
			Enable:     uinfo.Enable,
			CreateTime: uinfo.CreateTime,
			UpdateTime: uinfo.UpdateTime,
			Gender:     datum.Gender,
			Avatar:     datum.Avatar,
			Address:    datum.Address,
			Email:      datum.Email,
			Roles:      rols,
		})
	}
	Success(c, data)
}

func (user) Profile(c *gin.Context) {
	var params req.PatchProfileUserReq
	err := c.BindJSON(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	err = db.DB.Model(entity.Profile{}).Where("id=?", params.Id).Updates(entity.Profile{
		Gender:   params.Gender,
		Address:  params.Address,
		Email:    params.Email,
		NickName: params.NickName,
	}).Error
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	Success(c, err)
}
func (user) Update(c *gin.Context) {
	var params req.PatchUserReq
	err := c.BindJSON(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	orm := db.DB.Model(entity.User{}).Where("id=?", params.Id)
	if params.Password != nil {
		orm.Update("password", fmt.Sprintf("%x", md5.Sum([]byte(*params.Password))))
	}
	if params.Enable != nil {
		orm.Update("enable", *params.Enable)
	}
	if params.Username != nil {
		orm.Update("username", *params.Username)
		db.DB.Model(entity.Profile{}).Where("userId=?", params.Id).Update("nickName", *params.Username)
	}
	if params.RoleIds != nil {
		db.DB.Where("userId=?", params.Id).Delete(entity.UserRolesRole{})
		if len(*params.RoleIds) > 0 {
			for _, i2 := range *params.RoleIds {
				db.DB.Model(entity.UserRolesRole{}).Create(&entity.UserRolesRole{
					UserId: params.Id,
					RoleId: i2,
				})
			}
		}
	}

	Success(c, err)
}

func (user) Add(c *gin.Context) {
	var params req.AddUserReq
	err := c.Bind(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		var newUser = entity.User{
			Username:   params.Username,
			Password:   fmt.Sprintf("%x", md5.Sum([]byte(params.Password))),
			Enable:     params.Enable,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		err = tx.Create(&newUser).Error
		if err != nil {
			return err
		}
		tx.Create(&entity.Profile{
			UserId:   newUser.ID,
			NickName: newUser.Username,
		})
		for _, id := range params.RoleIds {
			tx.Create(&entity.UserRolesRole{
				UserId: newUser.ID,
				RoleId: id,
			})
		}
		return nil
	})
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	Success(c, "")
}
func (user) Delete(c *gin.Context) {
	uid := c.Param("id")
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		tx.Where("id =?", uid).Delete(&entity.User{})
		tx.Where("userId =?", uid).Delete(&entity.UserRolesRole{})
		tx.Where("userId =?", uid).Delete(&entity.Profile{})
		return nil
	})
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	Success(c, "")
}
