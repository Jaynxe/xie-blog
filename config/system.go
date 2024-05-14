package config

type System struct {
	Host string `json:"host,omitempty" yaml:"host"`
	Port int   `json:"port,omitempty" yaml:"port"`
	Env  string `json:"env,omitempty" yaml:"env"`
}
