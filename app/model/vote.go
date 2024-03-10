package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func GetVotes() []Vote {
	ret := make([]Vote, 0)
	err := Conn.Table("vote").Find(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}

func GetVote(id int64) VoteWithOpt {
	var ret Vote
	err := Conn.Table("vote").Where("id = ?", id).First(&ret).Find(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	opt := make([]VoteOpt, 0)
	err = Conn.Table("vote_opt").Where("vote_id = ?", id).Find(&opt).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return VoteWithOpt{
		Vote: ret,
		Opt:  opt,
	}
}

func DoVote(userId int64, voteId int64, optIds []int64) bool {
	var ret Vote
	err := Conn.Table("vote").Where("id = ?", voteId).First(&ret).Find(&ret).Error
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	for _, value := range optIds {
		if err := Conn.Table("vote_opt").Where("vote_id = ?", value).Update("count", gorm.Expr("count + ?", 1)).Error; err != nil {
			fmt.Printf("err:%s", err.Error())
		}

		user := VoteOptUser{
			VoteId:     voteId,
			UserId:     userId,
			VoteOptId:  value,
			CreateTime: time.Now(),
		}
		_ = Conn.Create(&user).Error
	}

	return true
}
