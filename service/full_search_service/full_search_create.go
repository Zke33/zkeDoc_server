package full_search_service

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"gvd_server/models"
)

// FullSearchCreate 添加
func FullSearchCreate(doc models.DocModel) {
	if global.ESClient == nil {
		return
	}
	searchDataList := MarkdownParse(doc.ID, doc.Title, doc.Content)
	bulk := global.ESClient.Bulk().Index(models.FullTextModel{}.Index()).Refresh("true")
	for _, model := range searchDataList {
		req := elastic.NewBulkCreateRequest().Doc(models.FullTextModel{
			DocID: doc.ID,
			Title: model.Title,
			Body:  model.Body,
			Slug:  model.Slug,
		})
		bulk.Add(req)
	}
	res, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Errorf("%#v 数据添加失败 err:%s", doc, err.Error())
		return
	}
	logrus.Infof("添加全文搜索记录 %d 条", len(res.Succeeded()))
}
