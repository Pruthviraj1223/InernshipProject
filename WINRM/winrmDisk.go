package WINRM

import (
	"encoding/json"
	"fmt"
	"github.com/masterzen/winrm"
	"strings"
)

func standardizeSpaces(s string) string {

	return strings.Join(strings.Fields(s), " ")

}

func WinrmDisk(data map[string]interface{}) {

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

	a := "aa"

	var disk string

	commandForDisk := "Get-WmiObject win32_logicaldisk | Foreach-Object {$_.DeviceId,$_.Freespace,$_.Size -join \" \"}" //disksize

	disk, _, _, err = client.RunPSWithString(commandForDisk, a)

	var diskList []map[string]string

	diskStringArray := strings.Split(disk, "\n")

	for _, v := range diskStringArray {

		diskEachWord := strings.Split(standardizeSpaces(v), " ")

		if len(diskEachWord) == 0 {

			break

		}

		if len(diskEachWord) == 3 {

			temp := map[string]string{

				"disk.name": diskEachWord[0],

				"disk.free": diskEachWord[1],

				"disk.size": diskEachWord[2],
			}

			diskList = append(diskList, temp)

		}
		if len(diskEachWord) == 1 {

			temp := map[string]string{

				"disk.name": diskEachWord[0],

				"disk.free": "0",

				"disk.size": "0",
			}

			diskList = append(diskList, temp)

		}

	}

	var diskMap = make(map[string]interface{})

	diskMap["disk"] = diskList

	bytes, _ := json.MarshalIndent(diskMap, " ", " ")

	fmt.Println(string(bytes))

}
