package admin

import (
	"beego_admin/models"
	"github.com/astaxie/beego/logs"
)

type Message struct {
	models.Model
	UserId int
	Mobile string
	Content string
}

func (Message)TableName() string {
	return "company_message"
}

// 留言板列表
func (message *Message)List(page, pageSize int) (messages []Message, totalCount int64) {
	models.DB.Model(&messages).Count(&totalCount)
	err := models.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&messages).Error
	if err != nil{
		logs.Error("查询留言列表报错", err)
		return nil, 0
	}
	return
}
