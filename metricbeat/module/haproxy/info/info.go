package info

import (
	"fmt"

	"github.com/ashilokhvostov/beats/libbeat/common"
	"github.com/ashilokhvostov/beats/libbeat/logp"
	"github.com/ashilokhvostov/beats/metricbeat/mb"
	"github.com/ashilokhvostov/beats/metricbeat/module/haproxy"
)

const (
	// defaultSocket is the default path to the unix socket tfor stats on haproxy.
	statsMethod = "info"
	defaultAddr = "unix:///var/lib/haproxy/stats"
)

var (
	debugf = logp.MakeDebug("haproxy-info")
)

// init registers the MetricSet with the central registry.
// The New method will be called after the setup of the module and before starting to fetch data
func init() {
	if err := mb.Registry.AddMetricSet("haproxy", "info", New); err != nil {
		panic(err)
	}
}

// MetricSet type defines all fields of the MetricSet
// As a minimum it must inherit the mb.BaseMetricSet fields, but can be extended with
// additional entries. These variables can be used to persist data or configuration between
// multiple fetch calls.
type MetricSet struct {
	mb.BaseMetricSet
	statsAddr string
	counter   int
}

// New create a new instance of the MetricSet
// Part of new is also setting up the configuration by processing additional
// configuration entries if needed.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	logp.Warn("EXPERIMENTAL: The haproxy info metricset is experimental")

	return &MetricSet{
		BaseMetricSet: base,
		statsAddr:     base.Host(),
		counter:       1,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right format
// It returns the event which is then forward to the output. In case of an error, a
// descriptive error must be returned.
func (m *MetricSet) Fetch() (common.MapStr, error) {

	hapc, err := haproxy.NewHaproxyClient(m.statsAddr)
	if err != nil {
		return nil, fmt.Errorf("HAProxy Client error: %s", err)
	}

	res, err := hapc.GetInfo()

	if err != nil {
		return nil, fmt.Errorf("HAProxy Client error fetching %s: %s", statsMethod, err)
	}
	m.counter++

	mappedEvent, err := eventMapping(res)
	if err != nil {
		return nil, err
	}
	return mappedEvent, nil

}
