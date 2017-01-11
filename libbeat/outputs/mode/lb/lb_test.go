package lb

import (
	"testing"

	"github.com/ashilokhvostov/beats/libbeat/common"
	"github.com/ashilokhvostov/beats/libbeat/logp"
	"github.com/ashilokhvostov/beats/libbeat/outputs"
)

var (
	testNoOpts     = outputs.Options{}
	testGuaranteed = outputs.Options{Guaranteed: true}

	testEvent = common.MapStr{
		"msg": "hello world",
	}
)

func enableLogging(selectors []string) {
	if testing.Verbose() {
		logp.LogInit(logp.LOG_DEBUG, "", false, true, selectors)
	}
}
