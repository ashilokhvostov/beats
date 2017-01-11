package main

import (
	"os"

	"github.com/ashilokhvostov/beats/libbeat/beat"
	"github.com/ashilokhvostov/beats/packetbeat/beater"

	// import support protocol modules
	_ "github.com/ashilokhvostov/beats/packetbeat/protos/amqp"
	_ "github.com/ashilokhvostov/beats/packetbeat/protos/cassandra"
	_ "github.com/ashilokhvostov/beats/packetbeat/protos/dns"
	_ "github.com/ashilokhvostov/beats/packetbeat/protos/http"
	_ "github.com/ashilokhvostov/beats/packetbeat/protos/memcache"
	_ "github.com/ashilokhvostov/beats/packetbeat/protos/mongodb"
	_ "github.com/ashilokhvostov/beats/packetbeat/protos/mysql"
	_ "github.com/ashilokhvostov/beats/packetbeat/protos/nfs"
	_ "github.com/ashilokhvostov/beats/packetbeat/protos/pgsql"
	_ "github.com/ashilokhvostov/beats/packetbeat/protos/redis"
	_ "github.com/ashilokhvostov/beats/packetbeat/protos/thrift"
)

var Name = "packetbeat"

// Setups and Runs Packetbeat
func main() {
	if err := beat.Run(Name, "", beater.New); err != nil {
		os.Exit(1)
	}
}
