package models

import (
	"myapp/internal/utils"

	"github.com/shopspring/decimal"
)

type UserInfo struct {
	Nickname string `json: "nickname"` //昵称
	Age      int    `json: "age"`      //年龄
	Gender   string `json: "gender"`   //性别
	Address  string `json: "address"`  //地址
}

type User struct {
	// DateStruct    DateStruct       `gorm:"embedded" json:"date"`                                   //通用date
	Uid           int                  `gorm:"primary_key;column:uid" json:"uid"`                      //uid
	Account       string               `gorm:"unique;not null;column:account" json:"account"`          //账号
	PhoneNumber   string               `gorm:"unique;not null;column:phone_number" json:"phoneNumber"` //手机号码
	Passwd        string               `gorm:"not null;column:passwd" json:"-"`                        //密码
	Roles         string               `gorm:"not null;column:roles"  json:"roles"`                    //角色
	Name          string               `gorm:"column:name"  json:"name"`                               //姓名
	Email         string               `gorm:"column:email" json:"email"`                              //邮箱
	Status        int8                 `gorm:"not null;column:status" json:"status"`                   //状态1启用2停止3冻结
	TotalMoney    decimal.Decimal      `gorm:"column:total_money" json:"totalMoney"`                   //总金额
	TotalIntegral int                  `gorm:"column:total_integral" json:"totalIntegral"`             //总积分
	LastTime      *utils.LocalDate     `gorm:"column:last_time" json:"lastTime"`                       //最后登录时间
	CreateTime    *utils.LocalDateTime `gorm:"column:create_time" json:"createTime"`                   //创建时间
	UpdateTime    *utils.LocalDateTime `gorm:"column:update_time" json:"updateTime"`                   //更新时间
	Salt          string               `gorm:"column:salt" json:"-"`                                   //盐
	Remark        string               `gorm:"column:remark" json:"remark"`                            //备注
	UserInfo      utils.JSON           `gorm:"column:user_info;type:json;" json:"userInfo"`            //用户详情json

	// LastTime      *utils.LocalDate `gorm:"column:last_time" json:"lastTime"`                       //最后登录时间
	//UserInfo      UserInfo         `gorm:embedded;column:user_info" json:"userInfo"`               //用户详情
}

func (u User) TableName() string {
	return "u_user"
}
