package {{ cookiecutter.module }}

import (
	"github.com/ashilokhvostov/beats/packetbeat/config"
	"github.com/ashilokhvostov/beats/packetbeat/protos"
)

type {{ cookiecutter.module }}Config struct {
	config.ProtocolCommon `config:",inline"`
}

var (
	defaultConfig = {{ cookiecutter.module }}Config{
		ProtocolCommon: config.ProtocolCommon{
			TransactionTimeout: protos.DefaultTransactionExpiration,
		},
	}
)

func (c *{{ cookiecutter.module }}Config) Validate() error {
	return nil
}
