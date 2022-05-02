package main

import (
	"InernshipProject/SNMP"
	"InernshipProject/SSH"
	"InernshipProject/WINRM"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	encoded := os.Args[1]

	jsonStr, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})

	err = json.Unmarshal(jsonStr, &data)

	if err != nil {

		fmt.Println("error ", err.Error())

	}

	if data["metricType"] == "linux" {

		if data["category"] == "discovery" {

			var ans = SSH.SshDiscovery(data)

			fmt.Println(ans)

		} else if data["category"] == "polling" {

			if data["counter"] == "disk" {
				SSH.SshDisk(data)
			} else if data["counter"] == "CPU" {
				SSH.SshCPU(data)
			} else if data["counter"] == "Memory" {
				SSH.SshMemory(data)
			} else if data["counter"] == "Process" {
				SSH.SshProcess(data)
			} else if data["counter"] == "SystemInfo" {
				SSH.SshSystem(data)
			}

		}
	} else if data["metricType"] == "windows" {

		if data["category"] == "discovery" {

			var ans = WINRM.WinrmDiscovery(data)

			fmt.Println(ans)

		} else if data["category"] == "polling" {

		}
	} else if data["metricType"] == "networking" {

		if data["category"] == "discovery" {

			var ans = SNMP.SnmpDiscovery(data)

			fmt.Println(ans)

		} else if data["category"] == "polling" {

			if data["counter"] == "systemInfo" {

				SNMP.SnmpSystem(data)

			}
		}
	}

}
