package models

import (
	"gvd_server/global"
	"sort"
	"strings"
)

type DocModel struct {
	Model
	Title           string      `gorm:"comment:文档标题;column:title;size:32" json:"title"`
	Content         string      `gorm:"comment:文档内容;column:content" json:"-"`
	DiggCount       int         `gorm:"comment:点赞量;column:diggCount" json:"diggCount"`
	LookCount       int         `gorm:"comment:浏览量;column:lookCount" json:"lookCount"`
	Key             string      `gorm:"comment:key;column:key;not null;unique" json:"key"`
	ParentID        *uint       `gorm:"comment:父文档id;column:parent_id" json:"parentID"`
	ParentModel     *DocModel   `gorm:"foreignKey:ParentID" json:"-"` // 父文档
	Child           []*DocModel `gorm:"foreignKey:ParentID" json:"-"` // 它会有子孙文档
	FreeContent     string      `gorm:"comment:预览部分;column:freeContent" json:"freeContent"`
	UserCollDocList []UserModel `gorm:"many2many:user_coll_doc_models;joinForeignKey:DocID;JoinReferences:UserID" json:"-"`
}

// FindAllParentDocList 找一个文档的所有父文档
func FindAllParentDocList(doc DocModel, docList *[]DocModel) {
	// 不管谁来，先把自己放进去
	*docList = append(*docList, doc)
	if doc.ParentID != nil {
		// 说明有父文档
		var parentDoc DocModel
		global.DB.Take(&parentDoc, *doc.ParentID)
		FindAllParentDocList(parentDoc, docList)
	}
}

// FindAllSubDocList 找一个文档的所有子文档
func FindAllSubDocList(doc DocModel) (docList []DocModel) {
	global.DB.Preload("Child").Take(&doc)
	for _, model := range doc.Child {
		docList = append(docList, *model)
		docList = append(docList, FindAllSubDocList(*model)...)
	}
	return
}

// DocTree 返回文档树
func DocTree(parentID *uint) (docList []*DocModel) {
	var query = global.DB.Where("")
	if parentID == nil {
		// id=nil，先找根文档
		query.Where("parent_id is null")
	} else {
		// 找谁的父文档id是parentID的
		query.Where("parent_id = ?", *parentID)
	}
	global.DB.Preload("Child").Where(query).Find(&docList)
	for _, model := range docList {
		subDocs := DocTree(&model.ID)
		model.Child = subDocs
	}
	return
}

// SortDocByPotCount 按照点的个数进行排序  返回最小的那个元素点的个数
func SortDocByPotCount(docList []*DocModel) (minCount int) {
	if len(docList) == 0 {
		return
	}
	sort.Slice(docList, func(i, j int) bool {
		count1 := GetByPotCount(docList[i])
		count2 := GetByPotCount(docList[j])
		if count1 == count2 {
			// 点的个数相同，按照id升序
			return docList[i].ID < docList[j].ID
		}
		return count1 < count2
	})
	return GetByPotCount(docList[0])
}

// GetByPotCount 获取文档点的个数
func GetByPotCount(doc *DocModel) int {
	return strings.Count(doc.Key, ".")
}

// TreeByOneDimensional 树的一维化
func TreeByOneDimensional(docList []*DocModel) (list []*DocModel) {
	for _, model := range docList {
		list = append(list, model)
		list = append(list, TreeByOneDimensional(model.Child)...)
	}
	return
}
