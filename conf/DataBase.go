package conf

// DataBase struct
type DataBase struct {
	DbEngine      string `yaml:"DbEngine"`
	CrawlerMaster string `yaml:"CrawlerMaster"`
	CrawlerSlave  string `yaml:"CrawlerSlave"`
}
