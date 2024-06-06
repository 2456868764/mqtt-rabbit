package api

import (
	"bifromq_engine/pkg/model/req"
	"strconv"

	"bifromq_engine/pkg/db"
	"bifromq_engine/pkg/model/entity"
	"bifromq_engine/pkg/model/resp"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Role = &role{}

type role struct {
}

func (role) PermissionsTree(c *gin.Context) {
	var uid, _ = c.Get("uid")

	var adminRole int64
	db.DB.Model(entity.UserRolesRole{}).Where("userId=? and roleId=1", uid).Count(&adminRole)
	orm := db.DB.Model(entity.Permission{}).Where("parentId is NULL").Order("`order` Asc")

	if adminRole == 0 {
		uroleIdList := db.DB.Model(entity.UserRolesRole{}).Where("userId=?", uid).Select("roleId")
		rpermisId := db.DB.Model(entity.RolePermissionsPermission{}).Where("roleId in(?)", uroleIdList).Select("permissionId")
		orm = orm.Where("id in(?)", rpermisId)
	}

	var onePermissList []entity.Permission
	orm.Find(&onePermissList)

	for i, perm := range onePermissList {
		var twoPerissList []entity.Permission
		db.DB.Model(entity.Permission{}).Where("parentId = ?", perm.ID).Order("`order` Asc").Find(&twoPerissList)
		for i2, perm2 := range twoPerissList {
			var twoPerissList2 []entity.Permission
			db.DB.Model(entity.Permission{}).Where("parentId = ?", perm2.ID).Order("`order` Asc").Find(&twoPerissList2)
			twoPerissList[i2].Children = twoPerissList2
		}
		onePermissList[i].Children = twoPerissList
	}
	Success(c, onePermissList)
}

func (role) List(c *gin.Context) {
	var data = &resp.RoleListRes{}

	db.DB.Model(entity.Role{}).Find(&data)

	Success(c, data)
}
func (role) ListPage(c *gin.Context) {
	var data = &resp.RoleListPageRes{}
	var name = c.DefaultQuery("name", "")
	var enable = c.DefaultQuery("enable", "")
	var pageNoReq = c.DefaultQuery("pageNo", "1")
	var pageSizeReq = c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(pageNoReq)
	pageSize, _ := strconv.Atoi(pageSizeReq)
	orm := db.DB.Model(entity.Role{})
	if name != "" {
		orm = orm.Where("name like ?", "%"+name+"%")
	}
	if enable != "" {
		ena := false
		if enable == "1" {
			ena = true
		}
		orm = orm.Where("enable = ?", ena)
	}
	orm.Count(&data.Total)

	orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&data.PageData)
	for i, datum := range data.PageData {
		var perIdList []int64
		db.DB.Model(entity.RolePermissionsPermission{}).Where("roleId=?", datum.ID).Select("permissionId").Find(&perIdList)
		data.PageData[i].PermissionIds = perIdList
	}
	Success(c, data)
}
func (role) Update(c *gin.Context) {
	var params req.PatchRoleReq
	err := c.BindJSON(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	orm := db.DB.Model(entity.Role{}).Where("id=?", params.Id)
	if params.Name != nil {
		orm.Update("name", *params.Name)
	}
	if params.Enable != nil {
		orm.Update("enable", *params.Enable)
	}
	if params.Code != nil {
		orm.Update("code", *params.Code)
	}
	if params.PermissionIds != nil {
		db.DB.Where("roleId=?", params.Id).Delete(entity.RolePermissionsPermission{})
		if len(*params.PermissionIds) > 0 {
			for _, i2 := range *params.PermissionIds {
				db.DB.Model(entity.RolePermissionsPermission{}).Create(&entity.RolePermissionsPermission{
					PermissionId: i2,
					RoleId:       params.Id,
				})
			}
		}
	}
	Success(c, err)
}

func (role) Add(c *gin.Context) {
	var params req.AddRoleReq
	err := c.Bind(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		var record = entity.Role{
			Code:   params.Code,
			Name:   params.Name,
			Enable: params.Enable,
		}
		err = tx.Create(&record).Error
		if err != nil {
			return err
		}

		for _, id := range params.PermissionIds {
			tx.Create(&entity.RolePermissionsPermission{
				RoleId:       record.ID,
				PermissionId: id,
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

func (role) Delete(c *gin.Context) {
	uid := c.Param("id")
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		tx.Where("id =?", uid).Delete(&entity.Role{})
		tx.Where("roleId =?", uid).Delete(&entity.UserRolesRole{})
		tx.Where("roleId =?", uid).Delete(&entity.RolePermissionsPermission{})
		return nil
	})
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	Success(c, "")
}
func (role) AddUser(c *gin.Context) {
	var params req.PatchRoleOpeateUserReq
	err := c.BindJSON(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	uid, _ := strconv.Atoi(c.Param("id"))
	params.Id = uid
	db.DB.Where("userId in (?) and roleId = ?", params.UserIds, params.Id).Delete(entity.UserRolesRole{})
	for _, id := range params.UserIds {
		db.DB.Model(entity.UserRolesRole{}).Create(entity.UserRolesRole{
			UserId: id,
			RoleId: params.Id,
		})
	}
	Success(c, "")
}
func (role) RemoveUser(c *gin.Context) {
	var params req.PatchRoleOpeateUserReq
	err := c.BindJSON(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	uid, _ := strconv.Atoi(c.Param("id"))
	params.Id = uid
	db.DB.Where("userId in (?) and roleId = ?", params.UserIds, params.Id).Delete(entity.UserRolesRole{})
	Success(c, "")
}
