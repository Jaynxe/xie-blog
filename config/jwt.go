package config

type Jwt struct {
	Signingkey    string `yaml:"signing-key"`
	ExpiresTime   string `yaml:"expires-time"`
	BufferTime    string `yaml:"buffer-time"`
	Issuer        string `yaml:"issuer"`
	SigningMethod string `yaml:"signing-method"`
}
