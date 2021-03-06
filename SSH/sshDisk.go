package SSH

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strconv"
	"strings"
	"time"
)

func SshDisk(data map[string]interface{}) {
	sshUser := (data["name"]).(string)

	sshPassword := (data["password"]).(string)

	sshHost := (data["ip"]).(string)

	sshPort := int((data["port"]).(float64))

	config := &ssh.ClientConfig{

		Timeout: 10 * time.Second,

		User: sshUser,

		HostKeyCallback: ssh.InsecureIgnoreHostKey(),

		Config: ssh.Config{Ciphers: []string{

			"aes128-ctr", "aes192-ctr", "aes256-ctr",
		}},
	}

	config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}

	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)

	sshClient, err := ssh.Dial("tcp", addr, config)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer sshClient.Close()

	session, err := sshClient.NewSession()

	if err != nil {

	}

	diskMap := make(map[string]interface{})

	var diskList []map[string]interface{}

	diskData, err := session.Output("df")

	diskUtilizationString := string(diskData)

	diskStringArray := strings.Split(diskUtilizationString, "\n")

	count := 1

	for _, v := range diskStringArray {

		if count == 1 {

			count++

			continue
		}

		diskEachWorld := strings.Split(standardizeSpaces(v), " ")

		if len(diskEachWorld) <= 5 {

			continue
		}

		usePercentString := strings.Trim(diskEachWorld[4], "-%")

		usePercent, _ := strconv.Atoi(usePercentString)

		freePercent := 100 - usePercent

		usedBytes, _ := strconv.Atoi(diskEachWorld[2])

		available, _ := strconv.Atoi(diskEachWorld[3])

		temp := make(map[string]interface{})

		temp["disk.name"] = diskEachWorld[0]

		temp["disk.total.bytes"] = diskEachWorld[1]

		temp["disk.used.bytes"] = usedBytes

		temp["disk.available.bytes"] = available

		temp["disk.used.percent"] = usePercent

		temp["disk.free.percent"] = freePercent

		diskList = append(diskList, temp)

	}

	diskMap["Disk"] = diskList

	bytes, _ := json.MarshalIndent(diskMap, " ", " ")

	fmt.Println(string(bytes))

}
