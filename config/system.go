package config

import (
	"fmt"
	"strconv"
)

type System struct {
	Host string `json:"host,omitempty" yaml:"host"`
	Port int    `json:"port,omitempty" yaml:"port"`
	Env  string `json:"env,omitempty" yaml:"env"`
}

func (s *System) Addr() string {
	return fmt.Sprintf(s.Host + ":" + strconv.Itoa(s.Port))
}
