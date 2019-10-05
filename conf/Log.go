package conf

// Log ...
type Log struct {
	LogPrefix    string `yaml:"logprefix"`
	TimeFileName string `yaml:"timefilename"`
	SaveMaxAge   int    `yaml:"savemaxage"`
	RotationTime int    `yaml:"rotationtime"`
}
