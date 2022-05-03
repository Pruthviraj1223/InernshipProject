package SNMP

import (
	g "github.com/gosnmp/gosnmp"
	"time"
)

func SnmpDiscovery(data map[string]interface{}) string {

	host := data["ip"].(string)

	port := (data["port"]).(float64)

	community := (data["community"]).(string)

	params := &g.GoSNMP{

		Target: host,

		Port: uint16(port),

		Community: community,

		Version: g.Version2c,

		Timeout: time.Duration(2) * time.Second,
	}

	err := params.Connect()

	_, err = params.Get([]string{"1.3.6.1.2.1.1.5.0"})

	if err != nil {

		return err.Error()

	}

	defer params.Conn.Close()

	if err != nil {

		return err.Error()

	}

	return "true"
}
