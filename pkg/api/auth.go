package api

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"bifromq_engine/pkg/db"
	"bifromq_engine/pkg/errcode"
	"bifromq_engine/pkg/model/entity"
	"bifromq_engine/pkg/model/req"
	"bifromq_engine/pkg/model/resp"
	"bifromq_engine/pkg/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var Auth = &auth{}

type auth struct {
}

func (auth) Captcha(c *gin.Context) {
	svg, code := utils.GenerateSVG(80, 40)
	session := sessions.Default(c)
	session.Set("captch", code)
	session.Save()
	// 设置 Content-Type 为 "image/svg+xml"
	c.Header("Content-Type", "image/svg+xml; charset=utf-8")
	// 返回验证码
	c.Data(http.StatusOK, "image/svg+xml", svg)
}

func (auth) Login(c *gin.Context) {
	var params req.LoginReq
	err := c.Bind(&params)
	if err != nil {
		ErrorStatus(c, errcode.ParamsError)
		return
	}
	session := sessions.Default(c)
	if params.Captcha != session.Get("captch") {
		Error(c, 20001, "验证码不正确")
		return
	}
	var info *entity.User
	db.DB.Model(entity.User{}).
		Where("username=? ", params.Username).
		Where("password=?", fmt.Sprintf("%x", md5.Sum([]byte(params.Password)))).
		Find(&info)
	if info.ID == 0 {
		Error(c, 20001, "账号或密码不正确")
		return
	}
	Success(c, resp.LoginRes{
		AccessToken: utils.GenerateToken(info.ID),
	})
}

func (auth) password(c *gin.Context) {
	var params req.AuthPwReq
	err := c.Bind(&params)
	if err != nil {
		Error(c, 20001, err.Error())
		return
	}
	uid, _ := c.Get("uid")
	var oldCun int64
	db.DB.Model(entity.User{}).Where("id=? and password=?", uid, fmt.Sprintf("%x", md5.Sum([]byte(params.OldPassword))))
	if oldCun > 0 {
		db.DB.Model(entity.User{}).
			Where("id=? ", uid).
			Update("password", fmt.Sprintf("%x", md5.Sum([]byte(params.NewPassword))))
	}
	Success(c, true)
}
func (auth) Logout(c *gin.Context) {
	Success(c, true)
}
