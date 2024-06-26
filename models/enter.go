package models

import "time"

type Model struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

type Pagination struct {
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"limit"`
	Key   string `json:"key" form:"key"`
	Sort  string `json:"sort" form:"sort"`
}

type IDListRequest struct {
	IDList []uint `json:"idList" form:"idList" binding:"required" label:"id列表"`
}

type IDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}
