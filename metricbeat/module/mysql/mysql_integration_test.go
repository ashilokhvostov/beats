// +build integration

package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/ashilokhvostov/beats/metricbeat/mb/testing"
)

func TestNewDB(t *testing.T) {
	db, err := NewDB(GetMySQLEnvDSN())
	assert.NoError(t, err)

	err = db.Ping()
	assert.NoError(t, err)
}
