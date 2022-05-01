package SSH

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
	"time"
)

func SshSystem(data map[string]interface{}) {

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
		err.Error()
	}

	defer sshClient.Close()

	sesion, _ := sshClient.NewSession()

	res, _ := sesion.Output("vmstat")

	splittedStr := strings.Split(string(res), "\n")

	var systemMap = make(map[string]interface{})

	flag := 1
	for _, v := range splittedStr {
		if flag == 1 || flag == 2 || v == "" {
			flag++
			continue
		}

		output := strings.SplitN(standardizeSpaces(v), " ", 17)

		fmt.Println(output)

		systemMap["system.runningProcess"] = output[0]
		systemMap["system.blockingProcess"] = output[1]
		systemMap["system.context.switches"] = output[11]

	}

	fmt.Println(systemMap)

}
