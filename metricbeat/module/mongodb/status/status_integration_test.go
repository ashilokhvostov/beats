// +build integration

package status

import (
	"testing"

	"github.com/ashilokhvostov/beats/libbeat/common"
	mbtest "github.com/ashilokhvostov/beats/metricbeat/mb/testing"
	"github.com/ashilokhvostov/beats/metricbeat/module/mongodb"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	f := mbtest.NewEventFetcher(t, getConfig())
	event, err := f.Fetch()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), event)

	// Check event fields
	current := event["connections"].(common.MapStr)["current"].(int64)
	assert.True(t, current >= 0)

	available := event["connections"].(common.MapStr)["available"].(int64)
	assert.True(t, available > 0)

	commits := event["journaling"].(common.MapStr)["commits"].(int64)
	assert.True(t, commits >= 0)
}

func TestData(t *testing.T) {
	f := mbtest.NewEventFetcher(t, getConfig())
	err := mbtest.WriteEvent(f, t)
	if err != nil {
		t.Fatal("write", err)
	}
}

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "mongodb",
		"metricsets": []string{"status"},
		"hosts":      []string{mongodb.GetEnvHost() + ":" + mongodb.GetEnvPort()},
	}
}
