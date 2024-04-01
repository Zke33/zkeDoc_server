package models

type UserPwdDocModel struct {
	Model
	UserID uint `gorm:"column:userID" json:"userID"`
	DocID  uint `gorm:"column:docID" json:"docID"`
}
