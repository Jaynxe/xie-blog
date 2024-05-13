package config

type Logger struct {
	Level        string `json:"level,omitempty" yaml:"level"`
	Prefix       string `json:"prefix,omitempty" yaml:"prefix"`
	Director     string `json:"director,omitempty" yaml:"director"`
	ShowLine     bool   `json:"show_line,omitempty" yaml:"show-line"`
	LogInConsole bool   `json:"log_in_console,omitempty" yaml:"log-in-console"`
}
