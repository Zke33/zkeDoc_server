package config

type Site struct {
	Title    string `json:"title" yaml:"title"`
	Icon     string `json:"icon" yaml:"icon"`
	Abstract string `json:"abstract" yaml:"abstract"`
	IconHref string `json:"iconHref" yaml:"iconHref"`
	Href     string `json:"href" yaml:"href"`
	Footer   string `json:"footer" yaml:"footer"`
	Content  string `json:"content" yaml:"content"`
}
