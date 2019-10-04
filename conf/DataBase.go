package conf

// DataBase struct
type DataBase struct {
	DbEngine      string `yaml:"dbengine"`
	CrawlerMaster string `yaml:"crawlermaster"`
	CrawlerSlave  string `yaml:"crawlerslave"`
}
