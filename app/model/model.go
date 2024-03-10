package model

import "time"

type User struct {
	Id         int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT;comment:Primary Key" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(50)" json:"name"`
	Password   string    `gorm:"column:password;type:varchar(50)" json:"password"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:Create Time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;comment:Update Time" json:"update_time"`
}

func (m *User) TableName() string {
	return "user"
}

type Vote struct {
	Id         int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT;comment:Primary Key" json:"id"`
	Title      string    `gorm:"column:title;type:varchar(255);comment:Vote Title" json:"title"`
	Type       int       `gorm:"column:type;type:int(11);comment:Vote Type: 0-Single, 1-Muti" json:"type"`
	Status     int       `gorm:"column:status;type:int(11);comment:Vote Status: 0-Normal, 1-Overtime" json:"status"`
	Time       int64     `gorm:"column:time;type:bigint(20);comment:Vote Valid Time" json:"time"`
	UserId     int64     `gorm:"column:user_id;type:bigint(20);comment:Vote Person" json:"user_id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:Create Time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;comment:Update Time" json:"update_time"`
}

func (m *Vote) TableName() string {
	return "vote"
}

type VoteOpt struct {
	Id         int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT;comment:Primary Key" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(255)" json:"name"`
	VoteId     int64     `gorm:"column:vote_id;type:bigint(20)" json:"vote_id"`
	Count      int       `gorm:"column:count;type:int(11)" json:"count"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:Create Time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;comment:Update Time" json:"update_time"`
}

func (m *VoteOpt) TableName() string {
	return "vote_opt"
}

type VoteOptUser struct {
	Id         int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT;comment:Primary Key" json:"id"`
	UserId     int64     `gorm:"column:user_id;type:bigint(20)" json:"user_id"`
	VoteId     int64     `gorm:"column:vote_id;type:bigint(20)" json:"vote_id"`
	VoteOptId  int64     `gorm:"column:vote_opt_id;type:bigint(20)" json:"vote_opt_id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:Create Time" json:"create_time"`
}

func (m *VoteOptUser) TableName() string {
	return "vote_opt_user"
}

type VoteWithOpt struct {
	Vote Vote
	Opt  []VoteOpt
}
