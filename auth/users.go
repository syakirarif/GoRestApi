package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restapi/model"
)

type FormUser struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	ctx.String(http.StatusOK, "hello lagi %s", username)
}

func Profile(ctx *gin.Context) {
	username := ctx.Query("username")
	ctx.String(http.StatusOK, "Hello %s", username)
}

func RegisterForm(ctx *gin.Context) {
	var user FormUser
	data := ctx.ShouldBind(&user)
	if data != nil {
		ctx.String(http.StatusInternalServerError, "invalid gaes")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func Category(ctx *gin.Context) {
	username := ctx.Param("username")
	ctx.String(http.StatusOK, "Hello %s", username)
}

//from database
func ShowUser(ctx *gin.Context) {
	var user []model.User
	err := model.DB.Find(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": "Gagal",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"data":   user,
	})

}
