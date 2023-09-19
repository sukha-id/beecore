package log

type Config struct {
	Dir       string `yaml:"dir"`
	FileName  string `yaml:"file_name"`
	MaxSize   int    `yaml:"max_size"`
	LocalTime bool   `yaml:"local_time"`
	Compress  bool   `yaml:"compress"`
}
