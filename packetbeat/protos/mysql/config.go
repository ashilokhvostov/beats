package mysql

import (
	"github.com/ashilokhvostov/beats/packetbeat/config"
	"github.com/ashilokhvostov/beats/packetbeat/protos"
)

type mysqlConfig struct {
	config.ProtocolCommon `config:",inline"`
	MaxRowLength          int `config:"max_row_length"`
	MaxRows               int `config:"max_rows"`
}

var (
	defaultConfig = mysqlConfig{
		ProtocolCommon: config.ProtocolCommon{
			TransactionTimeout: protos.DefaultTransactionExpiration,
		},
		MaxRowLength: 1024,
		MaxRows:      10,
	}
)
