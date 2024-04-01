package models

type EsIndexInterFace interface {
	Index() string
	Mapping() string
}

type FullTextModel struct {
	DocID uint   `json:"docID"`
	ID    string `json:"ID"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Slug  string `json:"slug"`
}

func (FullTextModel) Index() string {
	return "gvd_server_full_text_index"
}

func (FullTextModel) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "body": {
        "type": "text"
      },
      "title": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "slug": {
        "type": "keyword"
      },
      "docID": {
        "type": "integer"
      }
    }
  }
}
`
}
