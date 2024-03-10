package logic

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lancerstadium/gote/app/model"
	"github.com/lancerstadium/gote/app/tools"
)

type User struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

func GetUser(user *User) map[string]any {
	ret := make(map[string]any)
	err := model.Conn.Table("user").Where("name = ?", user.Name).Find(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}

func GetLogin(context *gin.Context) {
	context.HTML(http.StatusOK, "login.tmpl", nil)
}

func PostLogin(context *gin.Context) {
	var user User
	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusOK, tools.Ecode{
			Message: err.Error(), // mysql info?
		})
	}
	ret := model.GetUser(user.Name)
	if ret.Id < 1 || ret.Password != user.Password {
		context.JSON(http.StatusOK, tools.Ecode{
			Message: "Username or password error!",
		})
	} else {
		context.SetCookie("name", user.Name, 3600, "/", "", true, false)
		context.JSON(http.StatusOK, tools.Ecode{
			Message: "Login success!",
		})
	}
}
