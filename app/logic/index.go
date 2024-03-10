package logic

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lancerstadium/gote/app/model"
)

func GetIndex(context *gin.Context) {
	ret := model.GetVotes()
	context.HTML(http.StatusOK, "index.tmpl", gin.H{"vote": ret})
}

func GetVote(context *gin.Context) {
	var id int64
	idStr := context.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	ret := model.GetVote(id)
	context.HTML(http.StatusOK, "vote.tmpl", gin.H{"vote": ret})
}

func PostVote(context *gin.Context) {

}
