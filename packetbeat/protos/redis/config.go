package redis

import (
	"github.com/ashilokhvostov/beats/packetbeat/config"
	"github.com/ashilokhvostov/beats/packetbeat/protos"
)

type redisConfig struct {
	config.ProtocolCommon `config:",inline"`
}

var (
	defaultConfig = redisConfig{
		ProtocolCommon: config.ProtocolCommon{
			TransactionTimeout: protos.DefaultTransactionExpiration,
		},
	}
)
