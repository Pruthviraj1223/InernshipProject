package WINRM

import (
	"encoding/json"
	"fmt"
	"github.com/masterzen/winrm"
	"strings"
)

func WinrmProcess(data map[string]interface{}) {

	host := (data["ip"]).(string)

	port := int((data["port"]).(float64))

	name := (data["name"]).(string)

	password := (data["password"]).(string)

	endpoint := winrm.NewEndpoint(host, port, false, false, nil, nil, nil, 0)

	client, err := winrm.NewClient(endpoint, name, password)

	_, err = client.CreateShell()

	if err != nil {

		err.Error()

	}

	commandForProces := "get-process"

	a := "aa"

	process, _, _, err := client.RunPSWithString(commandForProces, a)

	var processList []map[string]string

	processStringArray := strings.Split(process, "\n")

	flag := 1

	for _, v := range processStringArray {

		if flag <= 3 {

			flag++

			continue

		}

		processEachWorld := strings.SplitN(standardizeSpaces(v), " ", 8)

		if len(processEachWorld) <= 7 {

			break

		}

		temp := map[string]string{

			"process.name": processEachWorld[7],

			"process.id": processEachWorld[6],

			"process.cpu": processEachWorld[5],

			"process.virtualMemory": processEachWorld[4],

			"process.pageableMemory": processEachWorld[2],

			"process.handles": processEachWorld[0],
		}

		processList = append(processList, temp)

	}

	var processMap = make(map[string]interface{})

	processMap["process"] = processList

	bytes, _ := json.MarshalIndent(processMap, " ", " ")

	fmt.Println(string(bytes))
}
