package api

import (
	"strconv"

	"bifromq_engine/pkg/db"
	"bifromq_engine/pkg/model/entity"
	"bifromq_engine/pkg/model/req"
	"bifromq_engine/pkg/model/resp"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Permissions = &permissions{}

type permissions struct {
}

func (permissions) List(c *gin.Context) {
	var onePermissList = make([]entity.Permission, 0)
	db.DB.Model(entity.Permission{}).Where("parentId is NULL").Order("`order` Asc").Find(&onePermissList)
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

func (permissions) ListPage(c *gin.Context) {
	var data = &resp.RoleListPageRes{}
	var name = c.DefaultQuery("name", "")
	var pageNoReq = c.DefaultQuery("pageNo", "1")
	var pageSizeReq = c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(pageNoReq)
	pageSize, _ := strconv.Atoi(pageSizeReq)
	orm := db.DB.Model(entity.Role{})
	if name != "" {
		orm = orm.Where("name like ?", "%"+name+"%")
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
func (permissions) Add(c *gin.Context) {
	var params req.AddPermissionReq
	err := c.Bind(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}

	err = db.DB.Model(entity.Permission{}).Create(&entity.Permission{
		Name:      params.Name,
		Code:      params.Code,
		Type:      params.Type,
		ParentId:  params.ParentId, // insert value null
		Path:      params.Path,
		Icon:      params.Icon,
		Component: params.Component,
		Layout:    params.Layout,
		KeepAlive: IsTrue(params.KeepAlive),
		Show:      IsTrue(params.Show),
		Enable:    IsTrue(params.Enable),
		Order:     params.Order,
	}).Error
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	Success(c, "")
}
func (permissions) Delete(c *gin.Context) {
	id := c.Param("id")
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		tx.Where("id =?", id).Delete(&entity.Permission{})
		tx.Where("permissionId =?", id).Delete(&entity.RolePermissionsPermission{})
		return nil
	})
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	Success(c, "")
}
func (permissions) PatchPermission(c *gin.Context) {
	var params req.PatchPermissionReq
	err := c.BindJSON(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}

	err = db.DB.Model(entity.Permission{}).Where("id=?", params.Id).Updates(entity.Permission{
		Name:      params.Name,
		Code:      params.Code,
		Type:      params.Type,
		ParentId:  params.ParentId,
		Path:      params.Path,
		Icon:      params.Icon,
		Component: params.Component,
		Layout:    params.Layout,
		KeepAlive: params.KeepAlive,
		Method:    params.Component,
		Show:      params.Show,
		Enable:    params.Enable,
		Order:     params.Order,
	}).Error
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	Success(c, "")

}
func IsTrue(v bool) int {
	if v {
		return 1
	}
	return 0
}
