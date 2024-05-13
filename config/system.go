package config

type System struct {
	Host string `json:"host,omitempty" yaml:"host"`
	Port uint   `json:"port,omitempty" yaml:"port"`
	Env  string `json:"env,omitempty" yaml:"env"`
}
