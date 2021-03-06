package WINRM

import (
	"github.com/masterzen/winrm"
)

func WinrmDiscovery(data map[string]interface{}) string {

	host := (data["ip"]).(string)

	port := int((data["port"]).(float64))

	name := (data["name"]).(string)

	password := (data["password"]).(string)

	endpoint := winrm.NewEndpoint(host, port, false, false, nil, nil, nil, 0)

	client, err := winrm.NewClient(endpoint, name, password)

	_, err = client.CreateShell()

	if err != nil {
		return err.Error()
	}

	if err != nil {

		return err.Error()

	}

	return "true"
}
