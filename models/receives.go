package models

// 邮件接收表
type Receives struct {
	Id         uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT;comment:主键自增" json:"id"`
	From       string `gorm:"column:from;type:varchar(255);comment:邮箱发送者;NOT NULL" json:"from"`
	To         string `gorm:"column:to;type:varchar(255);comment:邮箱接收者;NOT NULL" json:"to"`
	Content    string `gorm:"column:content;type:text;comment:邮箱内容;NOT NULL" json:"content"`
	Code       string `gorm:"column:code;type:varchar(12);comment:验证码;NOT NULL" json:"code"`
	CreateTime uint   `gorm:"column:create_time;type:int(11) unsigned;comment:创建时间;NOT NULL;autoCreateTime" json:"create_time"`
}

func (m *Receives) TableName() string {
	return "receives"
}
