package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lancerstadium/gote/app/logic"
)

func New() {
	r := gin.Default()
	r.LoadHTMLGlob("app/view/*")

	// router
	// r.GET("/", logic.GetIndex)
	index := r.Group("")
	index.Use(checkUser)
	index.GET("/index", logic.GetIndex)
	index.GET("/vote", logic.GetVote)
	index.POST("/vote", logic.PostVote)

	r.GET("/", logic.GetIndex)
	r.GET("/login", logic.GetLogin)
	r.POST("/login", logic.PostLogin)

	if err := r.Run(":8080"); err != nil {
		panic("Gin: Start fail!")
	}
}

func checkUser(context *gin.Context) {
	name, err := context.Cookie("name")
	if err != nil || name == "" {
		context.Redirect(http.StatusFound, "/login")
	}
	context.Next()
}
