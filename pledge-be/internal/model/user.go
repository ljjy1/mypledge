package model

// User 用户表
type User struct {
	ID       uint64 `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Login    string `gorm:"column:login;type:varchar(255)" json:"login"`       // 用户账号
	Nike     string `gorm:"column:nike;type:varchar(255)" json:"nike"`         // 用户昵称
	Password string `gorm:"column:password;type:varchar(255)" json:"password"` // 加密后的密码
}

// TableName 返回用户表名
func (m *User) TableName() string {
	return "user"
}

// UserColumnNames 自定义查询字段白名单，防止 SQL 注入攻击
var UserColumnNames = map[string]bool{
	"id":       true,
	"login":    true,
	"nike":     true,
	"password": true,
}
