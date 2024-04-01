package models

// LoginModel 用户登录数据
type LoginModel struct {
	Model
	UserID    uint      `gorm:"column:userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"column:ip;size:20" json:"ip"` // 登录的ip
	NickName  string    `gorm:"column:nickName;size:42" json:"nickName"`
	UA        string    `gorm:"column:ua;size:256" json:"ua"` // ua
	Token     string    `gorm:"column:token;size:256" json:"token"`
	Device    string    `gorm:"column:device;size:256" json:"device"` // 登录设备
	Addr      string    `gorm:"column:addr;size:64" json:"addr"`
}
