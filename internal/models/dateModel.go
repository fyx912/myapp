package models

import "myapp/internal/utils"

type DateStruct struct {
	CreateTime *utils.LocalDateTime `gorm:"column:create_time" json:"createTime"` //创建时间
	UpdateTime *utils.LocalDateTime `gorm:"column:update_time" json:"updateTime"` //更新时间
}
