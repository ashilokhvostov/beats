package activity

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/ashilokhvostov/beats/libbeat/common"
	"github.com/ashilokhvostov/beats/metricbeat/mb"
	"github.com/ashilokhvostov/beats/metricbeat/module/postgresql"
)

// init registers the MetricSet with the central registry.
// The New method will be called after the setup of the module and before starting to fetch data
func init() {
	if err := mb.Registry.AddMetricSet("postgresql", "activity", New); err != nil {
		panic(err)
	}
}

// MetricSet type defines all fields of the Postgresql MetricSet
type MetricSet struct {
	mb.BaseMetricSet
	connectionString string
}

// New create a new instance of the MetricSet
// Part of new is also setting up the configuration by processing additional
// configuration entries if needed.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {

	config := struct {
		Hosts    []string `config:"hosts"    validate:"nonzero,required"`
		Username string   `config:"username"`
		Password string   `config:"password"`
	}{
		Username: "",
		Password: "",
	}

	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	url, err := postgresql.ParseURL(base.Host(), config.Username, config.Password,
		base.Module().Config().Timeout)
	if err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet:    base,
		connectionString: url,
	}, nil
}

// Fetch implements the data gathering and data conversion to the right format.
func (m *MetricSet) Fetch() ([]common.MapStr, error) {

	db, err := sql.Open("postgres", m.connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	results, err := postgresql.QueryStats(db, "SELECT * FROM pg_stat_activity")
	if err != nil {
		return nil, errors.Wrap(err, "QueryStats")
	}

	events := []common.MapStr{}
	for _, result := range results {
		events = append(events, eventMapping(result))
	}

	return events, nil
}
