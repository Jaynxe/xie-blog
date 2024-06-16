package config
/* 站点信息 */
type Site struct {
	Name        string `yaml:"name"`
	Tagline     string `yaml:"tagline"`
	Description string `yaml:"description"`
	URL         string `yaml:"url"`
	AuthorName  string `yaml:"author_name"`
	Bio         string `yaml:"bio"`
	Twitter     string `yaml:"twitter"`
	Facebook    string `yaml:"facebook"`
	Instagram   string `yaml:"instagram"`
	Email       string `yaml:"email"`
	Phone       string `yaml:"phone"`
	Address     string `yaml:"address"`
}