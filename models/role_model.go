package models

type RoleModel struct {
	Model
	Title    string      `gorm:"column:title;size:16;not null;comment:角色名称" json:"title"`                       // 角色的名称
	Pwd      string      `gorm:"column:pwd;size:64;comment:角色的密码" json:"-"`                                     // 角色密码
	IsSystem bool        `gorm:"column:isSystem;comment:是否是系统角色" json:"isSystem"`                               // 是否是系统角色
	DocsList []DocModel  `gorm:"many2many:role_doc_models;joinForeignKey:RoleID;JoinReferences:DocID" json:"-"` // 角色拥有的文档列表
	UserList []UserModel `gorm:"foreignKey:RoleID" json:"-"`
}
